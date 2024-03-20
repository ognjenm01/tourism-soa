package repo

import (
	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogStatusRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogStatusRepository) Create(status *model.BlogStatus) error {
	result := repo.DatabaseConnection.Create(status)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
