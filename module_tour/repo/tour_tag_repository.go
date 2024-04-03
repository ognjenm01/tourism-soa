package repo

import (
	"tour/model"

	"gorm.io/gorm"
)

type TourTagRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourTagRepository) GetAll() ([]model.TourTag, error) {
	tags := []model.TourTag{}
	result := repo.DatabaseConnection.Find(&tags)
	if result != nil {
		return tags, result.Error
	}
	return tags, nil
}

func (repo *TourTagRepository) Save(tag *model.TourTag) error {
	result := repo.DatabaseConnection.Save(tag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
