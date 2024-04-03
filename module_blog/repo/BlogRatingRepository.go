package repo

import (
	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRatingRepository) GetByBlog(blogId string) ([]model.BlogRating, error) {
	ratings := []model.BlogRating{}
	result := repo.DatabaseConnection.Where("blog_id=?", blogId).Find(&ratings)

	if result != nil {
		return ratings, result.Error
	}

	return ratings, nil
}
