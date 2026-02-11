package usecase

import (
	"blogging-platform-api/internal/entity"
	"context"
	"errors"
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

func (u *blogUsecase) Create(ctx context.Context, blog *entity.Blog) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	tagsMap := make(map[uint]struct{})
	for _, v := range blog.Tags {
		if _, found := tagsMap[v.ID]; found {
			return nil, errors.New("tags id must be unique")
		}

		tagsMap[v.ID] = struct{}{}
	}

	return u.repo.Create(ctx, blog)
}

func (u *blogUsecase) Update(ctx context.Context, id uint, updateBlog *entity.Blog) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	tagsMap := make(map[uint]struct{})
	for _, v := range updateBlog.Tags {
		if _, found := tagsMap[v.ID]; found {
			return nil, errors.New("tags id must be unique")
		}

		tagsMap[v.ID] = struct{}{}
	}

	return u.repo.Update(ctx, id, updateBlog)
}

func (u *blogUsecase) Delete(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.Delete(ctx, id)
}
