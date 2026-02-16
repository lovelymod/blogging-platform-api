package entity

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   *string        `json:"content"`
	Category  *string        `json:"category"`
	Tags      []Tag          `json:"tags" gorm:"many2many:blog_tags;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type BlogFilter struct {
	Title    string `form:"title"`
	Category string `form:"category"`
	Tags     []uint `form:"tags"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

type CreateBlogRequest struct {
	Title    string  `json:"title" binding:"required"`
	Content  *string `json:"content"`
	Category *string `json:"category"`
	Tags     []uint  `json:"tags"`
}

type UpdateBlogRequest struct {
	Title    string  `json:"title" binding:"required"`
	Content  *string `json:"content"`
	Category *string `json:"category"`
	Tags     []uint  `json:"tags"`
}

type BlogRepository interface {
	GetAll(ctx context.Context, filter *BlogFilter) ([]Blog, int64, error)
	GetByID(ctx context.Context, id uint) (*Blog, error)
	Create(ctx context.Context, blog *Blog) (*Blog, error)
	Update(ctx context.Context, id uint, blog *Blog) (*Blog, error)
	Delete(ctx context.Context, id uint) error
}
type BlogUsecase interface {
	GetAll(ctx context.Context, filter *BlogFilter) ([]Blog, int64, error)
	GetByID(ctx context.Context, id uint) (*Blog, error)
	Create(ctx context.Context, blog *Blog) (*Blog, error)
	Update(ctx context.Context, id uint, blog *Blog) (*Blog, error)
	Delete(ctx context.Context, id uint) error
}
type BlogHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
