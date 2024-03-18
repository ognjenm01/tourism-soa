package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type TouristPositionService struct {
	TouristPositionRepository *repo.TouristPositionRepository
}

func (service *TouristPositionService) CreateTouristPosition(touristPosition *model.TouristPosition) error {
	error := service.TouristPositionRepository.Create(touristPosition)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TouristPositionService) GetTouristPositionById(id string) (*model.TouristPosition, error) {
	touristPosition, error := service.TouristPositionRepository.GetById(id)
	if error != nil {
		//log.Fatalf("[DB] - No tour review in db!\n")
		return nil, error
	}
	return &touristPosition, nil
}

func (service *TouristPositionService) UpdateTouristPosition(touristPosition *model.TouristPosition) error {
	error := service.TouristPositionRepository.Update(touristPosition)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
