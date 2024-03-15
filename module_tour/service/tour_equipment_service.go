package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type TourEquipmentService struct {
	TourEquipmentRepository *repo.TourEquipmentRepository
}

func (service *TourEquipmentService) CreateTourEquipment(tourEquipment *model.TourEquipment) error {
	error := service.TourEquipmentRepository.Create(tourEquipment)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourEquipmentService) GetAllTourEquipment() (*[]model.TourEquipment, error) {
	tourEquipment, error := service.TourEquipmentRepository.GetAll()
	if error != nil {
		log.Fatalf("[DB] - No equipment in db!\n")
		return nil, error
	}
	return &tourEquipment, nil
}

func (service *TourEquipmentService) DeleteTourEquipment(tourEquipment *model.TourEquipment) error {
	error := service.TourEquipmentRepository.Delete(tourEquipment)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
