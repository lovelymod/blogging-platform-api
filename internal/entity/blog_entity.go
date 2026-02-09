package entity

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Category  string         `json:"category"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type BlogRepository interface {
	Create(ctx context.Context, blog *Blog) error
	GetAll(ctx context.Context) ([]Blog, error)
	Delete(ctx context.Context, id uint) error
}

type BlogUsecase interface {
	Create(ctx context.Context, blog *Blog) error
	GetAll(ctx context.Context) ([]Blog, error)
	Delete(ctx context.Context, id uint) error
}
