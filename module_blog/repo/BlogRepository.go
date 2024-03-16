package repo

import (
	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) GetById(id int) (model.Blog, error) {
	blog := model.Blog{}
	result := repo.DatabaseConnection.Preload("BlogRatings").First(&blog, "blog_id = ?", id)

	if result != nil {
		return blog, result.Error
	}

	return blog, nil
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

	for _, img := range blog.ImageLinks {
		repo.DatabaseConnection.Save(img)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
