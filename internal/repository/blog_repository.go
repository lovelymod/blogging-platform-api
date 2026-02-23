package repository

import (
	"blogging-platform-api/internal/entity"
	"context"
	"errors"
	"log"

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

func (repo *blogRepository) GetAll(ctx context.Context, filter *entity.BlogFilter) ([]entity.Blog, int64, error) {
	blogs := make([]entity.Blog, 0)
	var totalRows int64

	tx := repo.db.WithContext(ctx).Model(&entity.Blog{})

	// LIKE = Case sensitive, ILIKE In case sensitve
	if filter.Title != "" {
		tx.Where("Title ILIKE ?", "%"+filter.Title+"%")
	}

	if filter.Category != "" {
		tx.Where("Category ILIKE ?", "%"+filter.Category+"%")
	}

	if len(filter.Tags) > 0 {
		tx.Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").Where("blog_tags.tag_id IN ?", filter.Tags).Group("blogs.id")
	}

	if err := tx.Count(&totalRows).Error; err != nil {
		log.Println(err)
		return nil, 0, entity.ErrGlobalServerErr
	}

	if filter.Limit > 0 {
		page := filter.Page
		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * filter.Limit
		tx.Offset(offset).Limit(filter.Limit)
	}

	if err := tx.Order("ID asc").Preload("User").Preload("Tags").Find(&blogs).Error; err != nil {
		log.Println(err)
		return nil, 0, entity.ErrGlobalServerErr
	}

	return blogs, totalRows, nil
}

func (repo *blogRepository) GetByID(ctx context.Context, id uint) (*entity.Blog, error) {
	var blog entity.Blog

	err := repo.db.WithContext(ctx).
		Preload("Tags").
		Preload("User").
		First(&blog, id).Error

	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrGlobalNotFound
		}

		return nil, entity.ErrGlobalServerErr
	}

	return &blog, nil
}

func (repo *blogRepository) Create(ctx context.Context, blog *entity.Blog) (*entity.Blog, error) {
	if err := repo.db.WithContext(ctx).Create(blog).Error; err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	var createdBlog entity.Blog

	if err := repo.db.WithContext(ctx).Preload("Tags").First(&createdBlog, blog.ID).Error; err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	return &createdBlog, nil
}

func (repo *blogRepository) Update(ctx context.Context, id uint, blog *entity.Blog) (*entity.Blog, error) {
	var existingBlog entity.Blog

	if err := repo.db.WithContext(ctx).Preload("Tags").First(&existingBlog, id).Error; err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	if err := repo.db.WithContext(ctx).Model(&existingBlog).Updates(blog).Error; err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	if err := repo.db.WithContext(ctx).Model(&existingBlog).Association("Tags").Replace(blog.Tags); err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	var updatedBlog entity.Blog

	if err := repo.db.WithContext(ctx).Preload("Tags").First(&updatedBlog, existingBlog.ID).Error; err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	return &updatedBlog, nil
}

func (repo *blogRepository) Delete(ctx context.Context, id uint) error {
	result := repo.db.WithContext(ctx).Delete(&entity.Blog{ID: id})

	if result.Error != nil {
		log.Println(result.Error)
		return entity.ErrGlobalServerErr
	}

	if result.RowsAffected == 0 {
		return entity.ErrGlobalNotFound
	}

	return nil
}
