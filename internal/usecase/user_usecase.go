package usecase

import (
	"blogging-platform-api/internal/entity"
	"context"
	"time"
)

type userUsercase struct {
	repo    entity.UserRepository
	timeout time.Duration
}

func NewUserUsecase(repo entity.UserRepository, timeout time.Duration) entity.UserUsercase {
	return &userUsercase{
		repo:    repo,
		timeout: timeout,
	}
}

func (u *userUsercase) GetUserProfile(username string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	existUser, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return existUser, nil
}
