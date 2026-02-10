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

	if err := repo.db.WithContext(ctx).Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *blogRepository) GetByID(ctx context.Context, id uint) (*entity.Blog, error) {
	var blog entity.Blog

	if err := repo.db.WithContext(ctx).Where(&entity.Blog{ID: id}).First(&blog).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *blogRepository) Create(ctx context.Context, blog *entity.Blog) error {
	return repo.db.WithContext(ctx).Create(blog).Error
}

func (repo *blogRepository) Update(ctx context.Context, id uint, updateBlog *entity.UpdateBlogRequest) (*entity.Blog, error) {
	var blog entity.Blog

	if result := repo.db.WithContext(ctx).First(&blog, id); result.Error != nil {
		return nil, result.Error
	}

	updateData := entity.Blog{
		Title:    updateBlog.Title,
		Content:  updateBlog.Content,
		Category: updateBlog.Category,
	}

	if result := repo.db.WithContext(ctx).Model(&blog).Updates(&updateData); result.Error != nil {
		return nil, result.Error
	}

	return &blog, nil
}

func (repo *blogRepository) Delete(ctx context.Context, id uint) error {
	result := repo.db.WithContext(ctx).Delete(&entity.Blog{ID: id})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
