package repo

import (
	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
	blog := model.Blog{}
	result := repo.DatabaseConnection.Preload("BlogStatuses").Preload("BlogRatings").First(&blog, "id = ?", id)

	if result != nil {

		return blog, result.Error
	}

	return blog, nil
}

func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
	blogs := []model.Blog{}
	result := repo.DatabaseConnection.Preload("BlogStatuses").Preload("BlogRatings").Find(&blogs)

	if result != nil {
		return blogs, result.Error
	}

	return blogs, nil
}

func (repo *BlogRepository) Create(blog *model.Blog) error {
	result := repo.DatabaseConnection.Create(blog)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *BlogRepository) Update(blog *model.Blog) error {
	result := repo.DatabaseConnection.Model(blog).
		Updates(map[string]interface{}{
			"Description":  blog.Description,
			"SystemStatus": blog.SystemStatus,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *BlogRepository) Delete(id string) error {
	result := repo.DatabaseConnection.Delete(model.Blog{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
