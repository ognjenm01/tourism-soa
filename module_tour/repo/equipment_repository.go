package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) Create(equipment *model.Equipment) error {
	result := repo.DatabaseConnection.Create(equipment)
	if result != nil {
		return result.Error
	}
	return nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	equipment := []model.Equipment{}
	result := repo.DatabaseConnection.Find(&equipment)
	if result != nil {
		return equipment, result.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) GetById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return equipment, err
	}
	result := repo.DatabaseConnection.First(&equipment, "id = ?", pk)
	if result != nil {
		return equipment, result.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) GetEquipmentByTourId(id string) ([]model.Equipment, error) {
	equipment := []model.Equipment{}
	result := repo.DatabaseConnection.Raw("select tours.equipment.id, name, description from tours.equipment join tours.tour_equipments on tours.equipment.id = tours.tour_equipments.id where tours.tour_equipments.tour_id = ?", id).Scan(&equipment)
	if result.Error != nil {
		return equipment, result.Error
	}
	return equipment, result.Error
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) error {
	result := repo.DatabaseConnection.Save(equipment)
	if result != nil {
		return result.Error
	}
	return nil
}
