package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) GetById(id string) (model.Tour, error) {
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

func (repo *TourRepository) GetAll() ([]model.Tour, error) {
	tours := []model.Tour{}
	result := repo.DatabaseConnection.Preload("Tags").Preload("Keypoints").Find(&tours)
	if result != nil {
		return tours, result.Error
	}
	return tours, nil
}

func (repo *TourRepository) Create(tour *model.Tour) error {
	result := repo.DatabaseConnection.Create(tour)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TourRepository) Update(id int, tour *model.Tour) error {
	//FIXME Transcations
	//repo.DatabaseConnection.Begin()
	result := repo.DatabaseConnection.Model(tour).
		Updates(map[string]interface{}{
			"Name":          tour.Name,
			"Description":   tour.Description,
			"Price":         tour.Price,
			"Difficulty":    tour.Difficulty,
			"TransportType": tour.TransportType,
		})
	for _, tag := range tour.Tags {
		repo.DatabaseConnection.Save(tag)
	}
	if result.Error != nil {
		//repo.DatabaseConnection.Rollback()
		return result.Error
	}
	//repo.DatabaseConnection.Commit()
	return nil
}
