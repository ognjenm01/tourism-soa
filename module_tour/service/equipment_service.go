package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type EquipmentService struct {
	EquipmentRepository *repo.EquipmentRepository
}

func (service *EquipmentService) CreateEquipment(equipment *model.Equipment) error {
	error := service.EquipmentRepository.Create(equipment)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *EquipmentService) GetAllEquipment() (*[]model.Equipment, error) {
	equipment, error := service.EquipmentRepository.GetAll()
	if error != nil {
		log.Fatalf("[DB] - No equipment in db!\n")
		return nil, error
	}
	return &equipment, nil
}

func (service *EquipmentService) GetEquipmentById(id string) (*model.Equipment, error) {
	equipment, error := service.EquipmentRepository.GetById(id)
	if error != nil {
		log.Fatalf("[DB] - No equipment in db!\n")
		return nil, error
	}
	return &equipment, nil
}

func (service *EquipmentService) GetEquipmentByTourId(tourId string) (*[]model.Equipment, error) {
	equipment, error := service.EquipmentRepository.GetEquipmentByTourId(tourId)
	if error != nil {
		log.Fatalf("Failled getting equipment by tourid %s", error)
		return nil, error
	}

	return &equipment, error
}

func (service *EquipmentService) UpdateEquipment(equipment *model.Equipment) error {
	error := service.EquipmentRepository.Update(equipment)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
