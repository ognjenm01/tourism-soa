package repo

import (
	"strconv"
	"tour/model"

	"gorm.io/gorm"
)

type TourReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourReviewRepository) FindReviewById(id string) (model.TourReview, error) {
	tourReview := model.TourReview{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return tourReview, err
	}
	result := repo.DatabaseConnection.First(&tourReview, "id = ?", pk)
	if result != nil {
		return tourReview, result.Error
	}
	return tourReview, nil
}

func (repo *TourReviewRepository) FindAllReviews() ([]model.TourReview, error) {
	tourReviews := []model.TourReview{}
	result := repo.DatabaseConnection.Preload("ImageLinks").Find(&tourReviews)
	if result != nil {
		return tourReviews, result.Error
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) CreateReview(tourReview *model.TourReview) error {
	result := repo.DatabaseConnection.Create((tourReview))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TourReviewRepository) UpdateReview(tourReview *model.TourReview) error {
	result := repo.DatabaseConnection.Save(tourReview)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
