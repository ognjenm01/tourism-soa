package service

import (
	"log"
	"strconv"
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

func (service *TouristPositionService) GetAllTouristPosition() (*[]model.TouristPosition, error) {
	touristPositions, error := service.TouristPositionRepository.GetAll()
	if error != nil {
		log.Printf("[DB] - No tourist positions in db!\n")
		return &[]model.TouristPosition{}, error
	}
	return &touristPositions, nil
}

func (service *TouristPositionService) GetTouristPositionById(id string) (*model.TouristPosition, error) {
	touristPosition, error := service.TouristPositionRepository.GetById(id)
	if error != nil {
		//log.Fatalf("[DB] - No tour review in db!\n")
		return nil, error
	}
	return &touristPosition, nil
}

func (service *TouristPositionService) GetTouristPositionByUser(userId string) (*[]model.TouristPosition, error) {
	UserId, error := strconv.Atoi(userId)
	if error != nil {
		log.Printf("[DB] - invalid userid!")
		return nil, error
	}

	touristPositions, error := service.TouristPositionRepository.GetAll()
	var filteredTouristPositions []model.TouristPosition
	if error == nil {
		for _, touristPosition := range touristPositions {
			if touristPosition.UserId == UserId {
				filteredTouristPositions = append(filteredTouristPositions, touristPosition)
			}
		}
		return &filteredTouristPositions, nil
	}
	return &[]model.TouristPosition{}, error
}

func (service *TouristPositionService) UpdateTouristPosition(touristPosition *model.TouristPosition) error {
	error := service.TouristPositionRepository.Update(touristPosition)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
