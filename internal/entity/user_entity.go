package entity

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Email          string         `json:"email" gorm:"unique;not null"`
	HashedPassword string         `json:"-" gorm:"not null"`
	Phone          string         `json:"phone"`
	Avatar         string         `json:"avatar"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserRegisterReq struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"hashed_password"`
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
}

type UserUsecase interface {
	Register(req *UserRegisterReq) error
}

type UserHandler interface {
	Register(c *gin.Context)
}
