package repository

import (
	"blogging-platform-api/internal/entity"
	"context"
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

func (repo *userRepository) Register(ctx context.Context, user *entity.User) error {
	var count int64

	if err := repo.db.WithContext(ctx).Model(&entity.User{}).Where(&entity.User{Email: user.Email}).Count(&count).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	if count != 0 {
		return entity.ErrUserThisEmailIsAlreadyUsed
	}

	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	return nil
}
