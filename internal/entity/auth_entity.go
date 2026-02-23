package entity

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"userID" gorm:"not null;index"`
	User      *User     `json:"user"`
	Token     string    `json:"token" gorm:"not null;unique"`
	Jti       string    `json:"jti" gorm:"not null;unique"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	IsRevoked bool      `json:"isRevoked" gorm:"default:false"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type AuthRegisterReq struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Username       string `json:"username"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"hashed_password"`
}

type AuthLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResp struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	User         *User  `json:"user,omitempty"`
}

type AuthRepository interface {
	CreateUser(ctx context.Context, registerUser *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetRefreshToken(ctx context.Context, claim *jwt.RegisteredClaims) (*RefreshToken, error)
	CreateRefreshToken(ctx context.Context, rt *RefreshToken) error
	UpdateRefreshToken(ctx context.Context, rt *RefreshToken) error
}

type AuthUsecase interface {
	Register(req *AuthRegisterReq) error
	Login(req *AuthLoginReq) (*AuthLoginResp, error)
	Logout(rt string) error
	RefreshToken(rt string) (string, string, error)
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}
