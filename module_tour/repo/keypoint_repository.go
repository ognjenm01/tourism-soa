package repo

import (
	"tour/model"

	"gorm.io/gorm"
)

type KeypointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeypointRepository) GetAll() ([]model.Keypoint, error) {
	keypoints := []model.Keypoint{}
	result := repo.DatabaseConnection.Find(&keypoints)
	if result != nil {
		return keypoints, result.Error
	}
	return []model.Keypoint{}, nil
}

func (repo *KeypointRepository) Save(keypoint *model.Keypoint) error {
	result := repo.DatabaseConnection.Save(keypoint)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *KeypointRepository) Delete(id string) error {
	result := repo.DatabaseConnection.Delete(model.Keypoint{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
