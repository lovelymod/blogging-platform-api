package repository

import (
	"blogging-platform-api/internal/entity"
	"context"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) entity.BlogRepository {
	return &blogRepository{
		db: db,
	}
}

func (repo *blogRepository) GetAll(ctx context.Context) ([]entity.Blog, error) {
	blogs := make([]entity.Blog, 0)

	if err := repo.db.WithContext(ctx).Preload("Tags").Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *blogRepository) GetByID(ctx context.Context, id uint) (*entity.Blog, error) {
	var blog entity.Blog

	if err := repo.db.WithContext(ctx).Preload("Tags").First(&blog, id).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *blogRepository) Create(ctx context.Context, blog *entity.Blog) error {
	if err := repo.db.WithContext(ctx).Create(blog).Error; err != nil {
		return err
	}

	return repo.db.WithContext(ctx).Preload("Tags").First(blog, blog.ID).Error
}

func (repo *blogRepository) Update(ctx context.Context, id uint, blog *entity.Blog) error {
	var existingBlog entity.Blog

	if result := repo.db.WithContext(ctx).Preload("Tags").First(&existingBlog, id); result.Error != nil {
		return result.Error
	}

	if result := repo.db.WithContext(ctx).Model(&existingBlog).Updates(blog); result.Error != nil {
		return result.Error
	}

	if err := repo.db.WithContext(ctx).Model(&existingBlog).Association("Tags").Replace(blog.Tags); err != nil {
		return err
	}

	return repo.db.WithContext(ctx).Preload("Tags").First(blog, existingBlog.ID).Error
}

func (repo *blogRepository) Delete(ctx context.Context, id uint) error {
	result := repo.db.WithContext(ctx).Delete(&entity.Blog{ID: id})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
