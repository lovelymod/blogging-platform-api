package usecase

import (
	"blogging-platform-api/internal/entity"
	"context"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo     entity.UserRepository
	timeout  time.Duration
	hashCost int
}

func NewUserUsecase(repo entity.UserRepository, contextTimeout time.Duration, env *entity.Config) entity.UserUsecase {
	cost, _ := strconv.Atoi(env.HASH_COST)
	if cost < 12 {
		cost = 12
	}

	return &userUsecase{
		repo:     repo,
		timeout:  contextTimeout,
		hashCost: cost,
	}
}

func (u *userUsecase) Register(req *entity.UserRegisterReq) error {
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

	return u.repo.Register(ctx, user)
}
