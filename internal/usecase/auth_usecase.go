package usecase

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"context"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authUsercase struct {
	repo               entity.AuthRepository
	timeout            time.Duration
	hashCost           int
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthUsecase(repo entity.AuthRepository, contextTimeout time.Duration, config *entity.Config) entity.AuthUsecase {
	cost, _ := strconv.Atoi(config.HASH_COST)
	if cost < 12 {
		cost = 12
	}

	return &authUsercase{
		repo:               repo,
		timeout:            contextTimeout,
		hashCost:           cost,
		accessTokenSecret:  config.ACCESS_TOKEN_SECRET,
		refreshTokenSecret: config.REFRESH_TOKEN_SECRET,
	}
}

func (u *authUsercase) Register(req *entity.AuthRegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	hasdPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), u.hashCost)
	if err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	user := &entity.User{
		Email:          req.Email,
		HashedPassword: string(hasdPassword),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	if req.Username != "" {
		user.Username = req.Username
		user.DisplayName = req.Username
	} else {
		defaultUsername := "user" + strconv.FormatInt(time.Now().Unix(), 10)
		user.Username = defaultUsername
		user.DisplayName = defaultUsername
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *authUsercase) Login(req *entity.AuthLoginReq) (*entity.AuthLoginResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	existUser, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.HashedPassword), []byte(req.Password)); err != nil {
		log.Println(err)
		return nil, entity.ErrAuthWrongEmailOrPassword
	}

	_, at, err := utils.SignAccessToken(existUser, u.accessTokenSecret)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	rtClaims, rt, err := utils.SignRefreshToken(existUser, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	savedRT := &entity.RefreshToken{
		UserID:    existUser.ID,
		Token:     rt,
		ExpiresAt: rtClaims.ExpiresAt.Time,
		Jti:       rtClaims.ID,
	}

	if err := u.repo.CreateRefreshToken(ctx, savedRT); err != nil {
		return nil, err
	}

	return &entity.AuthLoginResp{
		User:         existUser,
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (u *authUsercase) Logout(rt string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	cliams, err := utils.ParseRefreshToken(rt, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return err
	}

	updatedRT := &entity.RefreshToken{
		Jti:       cliams.ID,
		IsRevoked: true,
	}

	return u.repo.UpdateRefreshToken(ctx, updatedRT)
}

func (u *authUsercase) RefreshToken(rt string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	// Check refresh token is valid or not
	oldRTClaims, err := utils.ParseRefreshToken(rt, u.refreshTokenSecret)
	if err != nil {
		return "", "", err
	}

	// Get refreshToken in db
	existRT, err := u.repo.GetRefreshToken(ctx, oldRTClaims)
	if err != nil {
		return "", "", err
	}

	if existRT.IsRevoked || time.Now().After(existRT.ExpiresAt) {
		return "", "", entity.ErrAuthTokenExpired
	}

	// Sign new accessToken
	_, newAT, err := utils.SignAccessToken(&existRT.User, u.accessTokenSecret)
	if err != nil {
		return "", "", entity.ErrGlobalServerErr
	}

	// Sign new refreshToken
	newRTClaims, newRT, err := utils.SignRefreshToken(&existRT.User, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return "", "", entity.ErrGlobalServerErr
	}

	// Save new refreshToken to db
	userID, err := strconv.ParseUint(newRTClaims.Subject, 10, 64)
	if err != nil {
		log.Println(err)
		return "", "", entity.ErrGlobalServerErr
	}

	savedRT := &entity.RefreshToken{
		UserID:    uint(userID),
		Token:     newRT,
		ExpiresAt: newRTClaims.ExpiresAt.Time,
		Jti:       newRTClaims.ID,
	}

	if err := u.repo.CreateRefreshToken(ctx, savedRT); err != nil {
		return "", "", err
	}

	// Revoke old refreshToken
	existRT.IsRevoked = true
	if err := u.repo.UpdateRefreshToken(ctx, existRT); err != nil {
		return "", "", err
	}

	return newAT, newRT, nil
}
