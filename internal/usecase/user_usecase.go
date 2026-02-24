package usecase

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/internal/provider"
	"context"
	"time"
)

type userUsercase struct {
	repo       entity.UserRepository
	s3Provider provider.S3Provider
	timeout    time.Duration
}

func NewUserUsecase(repo entity.UserRepository, s3Provider provider.S3Provider, timeout time.Duration) entity.UserUsercase {
	return &userUsercase{
		repo:       repo,
		timeout:    timeout,
		s3Provider: s3Provider,
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
