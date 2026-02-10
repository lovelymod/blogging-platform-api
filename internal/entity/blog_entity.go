package entity

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   *string        `json:"content"`
	Category  *string        `json:"category"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UpdateBlogRequest struct {
	Title    string  `json:"title" binding:"required"`
	Content  *string `json:"content"`
	Category *string `json:"category"`
}

type BlogRepository interface {
	GetAll(ctx context.Context) ([]Blog, error)
	GetByID(ctx context.Context, id uint) (*Blog, error)
	Create(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, id uint, updateBlog *UpdateBlogRequest) (*Blog, error)
	Delete(ctx context.Context, id uint) error
}

type BlogUsecase interface {
	GetAll(ctx context.Context) ([]Blog, error)
	GetByID(ctx context.Context, id uint) (*Blog, error)
	Create(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, id uint, updateBlog *UpdateBlogRequest) (*Blog, error)
	Delete(ctx context.Context, id uint) error
}
