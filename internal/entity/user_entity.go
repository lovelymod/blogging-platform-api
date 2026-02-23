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
	DisplayName    string         `json:"displayName"`
	Email          string         `json:"email" gorm:"unique;not null"`
	Username       string         `json:"username" gorm:"unique;not null"`
	HashedPassword string         `json:"-" gorm:"not null"`
	Phone          string         `json:"phone"`
	Avatar         string         `json:"avatar"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}
type UserUsercase interface {
	GetUserProfile(username string) (*User, error)
}
type UserHandler interface {
	GetUserProfile(c *gin.Context)
}
