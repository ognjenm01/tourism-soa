package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type TourProgressRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourProgressRepository) Create(tourProgress *model.TourProgress) error {
	result := repo.DatabaseConnection.Create((tourProgress))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TourProgressRepository) GetAll() ([]model.TourProgress, error) {
	tourProgress := []model.TourProgress{}
	//result := repo.DatabaseConnection.Preload("TouristPosition").Find(&tourProgress)
	result := repo.DatabaseConnection.Find(&tourProgress)
	if result != nil {
		return tourProgress, result.Error
	}
	return tourProgress, nil
}

func (repo *TourProgressRepository) GetById(id string) (model.TourProgress, error) {
	tourProgress := model.TourProgress{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return tourProgress, err
	}
	result := repo.DatabaseConnection.First(&tourProgress, "id = ?", pk)
	if result.Error != nil {
		return tourProgress, result.Error
	}
	return tourProgress, nil
}

func (repo *TourProgressRepository) Update(tourProgress *model.TourProgress) error {
	result := repo.DatabaseConnection.Save(tourProgress)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TourProgressRepository) Delete(id string) error {
	result := repo.DatabaseConnection.Delete(model.TourProgress{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
