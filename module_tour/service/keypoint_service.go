package service

import (
	"log"
	"strconv"
	"tour/model"
	"tour/repo"
)

type KeypointService struct {
	KeypointRepository *repo.KeypointRepository
}

func (service *KeypointService) GetAll() (*[]model.Keypoint, error) {
	keypoints, error := service.KeypointRepository.GetAll()
	if error != nil {
		log.Fatalf("[DB] - No tours in db!\n")
		return &[]model.Keypoint{}, error
	}
	return &keypoints, nil
}

func (service *KeypointService) GetKeypointsByTourId(tourid string) (*[]model.Keypoint, error) {
	keypoints, error := service.KeypointRepository.GetAll()
	if error != nil {
		log.Fatalf("[DB] - No tours in db!\n")
		return &[]model.Keypoint{}, error
	}

	filteredKeypoints := []model.Keypoint{}
	ID, error := strconv.Atoi(tourid)
	if error != nil {
		log.Fatalf("[DB] - No keypoints in db!\n")
		return &[]model.Keypoint{}, error
	}
	for _, keypoint := range keypoints {
		if keypoint.TourID == ID {
			filteredKeypoints = append(filteredKeypoints, keypoint)
		}
	}
	return &filteredKeypoints, nil
}

func (service *KeypointService) CreateKeypoint(keypoint *model.Keypoint) error {
	error := service.KeypointRepository.Save(keypoint)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *KeypointService) DeleteKeypoint(id string) error {
	error := service.KeypointRepository.Delete(id)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
