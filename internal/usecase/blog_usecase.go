package usecase

import (
	"blogging-platform-api/internal/entity"
	"context"
	"time"
)

type blogUsecase struct {
	repo           entity.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(repo entity.BlogRepository, timeout time.Duration) entity.BlogUsecase {
	return &blogUsecase{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (u *blogUsecase) GetAll(ctx context.Context) ([]entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.GetAll(ctx)
}

func (u *blogUsecase) GetByID(ctx context.Context, id uint) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.GetByID(ctx, id)
}

func (u *blogUsecase) Create(ctx context.Context, blog *entity.Blog) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.Create(ctx, blog)
}

func (u *blogUsecase) Update(ctx context.Context, id uint, updateBlog *entity.UpdateBlogRequest) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.Update(ctx, id, updateBlog)
}

func (u *blogUsecase) Delete(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.Delete(ctx, id)
}
