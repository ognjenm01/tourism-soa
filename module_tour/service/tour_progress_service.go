package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type TourProgressService struct {
	TourProgressRepository *repo.TourProgressRepository
}

func (service *TourProgressService) CreateTourProgress(tourProgress *model.TourProgress) error {
	error := service.TourProgressRepository.Create(tourProgress)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourProgressService) GetAllTourProgress() (*[]model.TourProgress, error) {
	tourProgress, error := service.TourProgressRepository.GetAll()
	if error != nil {
		//log.Fatalf("[DB] - No tour progress in db!\n")
		return nil, error
	}
	return &tourProgress, nil
}

func (service *TourProgressService) GetTourProgressById(id string) (*model.TourProgress, error) {
	tourProgress, error := service.TourProgressRepository.GetById(id)
	if error != nil {
		//log.Fatalf("[DB] - No tour review in db!\n")
		return nil, error
	}
	return &tourProgress, nil
}

func (service *TourProgressService) UpdateTourProgress(tourProgress *model.TourProgress) error {
	error := service.TourProgressRepository.Update(tourProgress)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourProgressService) DeleteTourProgress(id string) error {
	error := service.TourProgressRepository.Delete(id)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
