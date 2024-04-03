package service

import (
	"fmt"
	"log"
	"strconv"
	"tour/model"
	"tour/repo"
)

type TourService struct {
	TourRepository *repo.TourRepository
}

func (service *TourService) GetTourById(id string) (*model.Tour, error) {
	tour, error := service.TourRepository.GetById(id)
	if error != nil {
		log.Printf("[DB] - Entity with id %s not found\n", id)
		return &model.Tour{}, fmt.Errorf(fmt.Sprintf("[DB] - Entity with id %s not found\n", id))
	}
	return &tour, nil
}

func (service *TourService) GetToursByStatus(statuses []model.TourStatus) (*[]model.Tour, error) {
	tours, error := service.TourRepository.GetAll()
	var filteredTours []model.Tour
	if error == nil {
		for _, tour := range tours {
			if tour.ContainsStatus(statuses) {
				filteredTours = append(filteredTours, tour)
			}
		}
		return &filteredTours, nil
	}

	log.Printf("[DB] - %s\n", error.Error())
	return &[]model.Tour{}, error
}

func (service *TourService) GetToursByAuthor(authorId string) (*[]model.Tour, error) {
	AuthorId, error := strconv.Atoi(authorId)
	if error != nil {
		log.Printf("[DB] - invalid userid!")
		return nil, error
	}

	tours, error := service.TourRepository.GetAll()
	var filteredTours []model.Tour
	if error == nil {
		for _, tour := range tours {
			if tour.UserId == AuthorId {
				filteredTours = append(filteredTours, tour)
			}
		}
		return &filteredTours, nil
	}
	return &[]model.Tour{}, error
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, error := service.TourRepository.GetAll()
	if error != nil {
		log.Printf("[DB] - No tours in db!\n")
		return &[]model.Tour{}, error
	}
	return &tours, nil
}

func (service *TourService) CreateTour(tour *model.Tour) error {
	error := service.TourRepository.Create(tour)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourService) UpdateTour(id int, tour *model.Tour) error {
	error := service.TourRepository.Update(id, tour)
	if error != nil {
		log.Printf("[DB] - %s\n", error)
		return error
	}
	return nil
}
