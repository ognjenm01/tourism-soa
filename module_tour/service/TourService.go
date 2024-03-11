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
	tour, error := service.TourRepository.FindById(id)
	if error != nil {
		log.Fatalf("[DB] - Entity with id %s not found\n", id)
		return nil, fmt.Errorf(fmt.Sprintf("[DB] - Entity with id %s not found\n", id))
	}
	return &tour, nil
}

func (service *TourService) GetToursByStatus(statuses []model.TourStatus) (*[]model.Tour, error) {
	tours, error := service.TourRepository.FindAll()
	var filteredTours []model.Tour
	if error == nil {
		for _, tour := range tours {
			if tour.ContainsStatus(statuses) {
				filteredTours = append(filteredTours, tour)
			}
		}
		return &filteredTours, nil
	}
	return &filteredTours, error
}

func (service *TourService) GetToursByAuthor(authorId string) (*[]model.Tour, error) {
	AuthorId, error := strconv.Atoi(authorId)
	if error != nil {
		log.Fatalln("[DB] - invalid userid!")
		return nil, error
	}

	tours, error := service.TourRepository.FindAll()
	var filteredTours []model.Tour
	if error == nil {
		for _, tour := range tours {
			if tour.UserId == AuthorId {
				filteredTours = append(filteredTours, tour)
			}
		}
		return &filteredTours, nil
	}
	return &filteredTours, error
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, error := service.TourRepository.FindAll()
	if error != nil {
		log.Fatalf("[DB] - No tours in db!\n")
		return nil, error
	}
	return &tours, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	error := service.TourRepository.Create(tour)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	error := service.TourRepository.Save(tour)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
