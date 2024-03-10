package service

import (
	"fmt"
	"log"
	"tour/model"
	"tour/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) FindTour(id string) (*model.Tour, error) {
	tour, error := service.TourRepository.FindById(id)
	if error != nil {
		log.Fatalf("[DB] - Entity with id %s not found\n", id)
		return nil, fmt.Errorf(fmt.Sprintf("[DB] - Entity with id %s not found\n", id))
	}
	return &tour, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	error := service.TourRepository.Create(tour)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
