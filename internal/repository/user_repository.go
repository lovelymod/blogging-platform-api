package repository

import (
	"blogging-platform-api/internal/entity"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var existUser entity.User

	if err := repo.db.WithContext(ctx).Where(&entity.User{Username: username}).First(&existUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrGlobalNotFound
		}
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	return &existUser, nil
}
