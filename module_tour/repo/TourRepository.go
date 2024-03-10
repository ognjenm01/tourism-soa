package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	tour := model.Tour{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return tour, err
	}
	result := repo.DatabaseConnection.Preload("Tags").First(&tour, "id = ?", pk)
	if result != nil {
		return tour, result.Error
	}
	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) error {
	result := repo.DatabaseConnection.Create(tour)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
