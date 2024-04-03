package repo

import (
	"tour/model"

	"gorm.io/gorm"
)

type TourEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourEquipmentRepository) GetAll() ([]model.TourEquipment, error) {
	tourEquipment := []model.TourEquipment{}
	result := repo.DatabaseConnection.Find(&tourEquipment)
	if result != nil {
		return tourEquipment, result.Error
	}
	return tourEquipment, nil
}

func (repo *TourEquipmentRepository) Create(tourEquipment *model.TourEquipment) error {
	tourEquipments, error := repo.GetAll()
	if error != nil {
		return error
	}

	for _, te := range tourEquipments {
		if te.TourID == tourEquipment.TourID && te.EquipmentID == tourEquipment.EquipmentID {
			return nil
		}
	}

	result := repo.DatabaseConnection.Create(tourEquipment)
	if result != nil {
		return result.Error
	}
	return nil
}

func (repo *TourEquipmentRepository) Delete(tourEquipment *model.TourEquipment) error {
	result := repo.DatabaseConnection.Delete(model.TourEquipment{}, "tour_id = ? and equipment_id = ?", tourEquipment.TourID, tourEquipment.EquipmentID)
	if result != nil {
		return result.Error
	}
	return nil
}
