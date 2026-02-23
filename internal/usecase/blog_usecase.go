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

func (u *blogUsecase) GetAll(ctx context.Context, filter *entity.BlogFilter) ([]entity.BlogResp, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	blogs, totalRows, err := u.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var blogsResp []entity.BlogResp

	for _, blog := range blogs {
		blogsResp = append(blogsResp, entity.BlogResp{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			Category:  blog.Category,
			Tags:      blog.Tags,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
			Author: &entity.AuthorResp{
				DisplayName: blog.User.DisplayName,
				Username:    blog.User.Username,
				Avatar:      blog.User.Avatar,
			},
		})
	}

	return blogsResp, totalRows, nil
}

func (u *blogUsecase) GetByID(ctx context.Context, id uint) (*entity.BlogResp, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	blog, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	blogResp := &entity.BlogResp{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		Category:  blog.Category,
		Tags:      blog.Tags,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		Author: &entity.AuthorResp{
			DisplayName: blog.User.DisplayName,
			Username:    blog.User.Username,
			Avatar:      blog.User.Avatar,
		},
	}

	return blogResp, nil
}

func (u *blogUsecase) Create(ctx context.Context, blog *entity.Blog) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	tagsMap := make(map[uint]struct{})
	for _, v := range blog.Tags {
		if _, found := tagsMap[v.ID]; found {
			return nil, entity.ErrBlogTagMustBeUnique
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
			return nil, entity.ErrBlogTagMustBeUnique
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
