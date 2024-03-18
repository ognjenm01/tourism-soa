package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type TouristPositionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TouristPositionRepository) Create(touristPosition *model.TouristPosition) error {
	result := repo.DatabaseConnection.Create((touristPosition))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TouristPositionRepository) GetById(id string) (model.TouristPosition, error) {
	touristPosition := model.TouristPosition{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return touristPosition, err
	}
	result := repo.DatabaseConnection.First(&touristPosition, "id = ?", pk)
	if result.Error != nil {
		return touristPosition, result.Error
	}
	return touristPosition, nil
}

func (repo *TouristPositionRepository) Update(touristPosition *model.TouristPosition) error {
	result := repo.DatabaseConnection.Save(touristPosition)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
