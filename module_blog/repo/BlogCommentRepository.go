package repo

import (
	"strconv"

	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogCommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogCommentRepository) GetById(id string) (model.BlogComment, error) {
	blogComment := model.BlogComment{}
	pk, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return blogComment, err
	}

	result := repo.DatabaseConnection.First(&blogComment, "id = ?", pk)

	if result != nil {
		return blogComment, result.Error
	}

	return blogComment, nil
}

func (repo *BlogCommentRepository) GetAll() ([]model.BlogComment, error) {
	blogComments := []model.BlogComment{}
	result := repo.DatabaseConnection.Find(&blogComments)

	if result != nil {
		return blogComments, result.Error
	}

	return blogComments, nil
}

func (repo *BlogCommentRepository) Create(blogComment *model.BlogComment) error {
	result := repo.DatabaseConnection.Create(blogComment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *BlogCommentRepository) Update(blogComment *model.BlogComment) error {
	result := repo.DatabaseConnection.Save(blogComment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *BlogCommentRepository) Delete(id string) error {
	result := repo.DatabaseConnection.Delete(model.BlogComment{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// IZMENI UPIT
func (repo *BlogCommentRepository) GetCommentsByBlogId(id string) ([]model.BlogComment, error) {
	blogComments := []model.BlogComment{}
	result := repo.DatabaseConnection.Raw("SELECT * FROM blog.blog_comments WHERE blog.blog_comments.blog_id = ?", id).Scan(&blogComments)
	if result.Error != nil {
		return blogComments, result.Error
	}
	return blogComments, result.Error
}
