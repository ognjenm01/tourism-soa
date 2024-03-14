package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type TourReviewService struct {
	TourReviewRepository *repo.TourReviewRepository
}

func (service *TourReviewService) FindReviewById(id string) (*model.TourReview, error) {
	tourReview, error := service.TourReviewRepository.FindReviewById(id)
	if error != nil {
		log.Fatalf("[DB] - No tour review in db!\n")
		return nil, error
	}
	return &tourReview, nil
}

func (service *TourReviewService) FindAllReviews() (*[]model.TourReview, error) {
	tourReviews, error := service.TourReviewRepository.FindAllReviews()
	if error != nil {
		log.Fatalf("[DB] - No tour reviews in db!\n")
		return nil, error
	}
	return &tourReviews, nil
}

func (service *TourReviewService) CreateReview(tourReview *model.TourReview) error {
	error := service.TourReviewRepository.CreateReview(tourReview)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourReviewService) UpdateReview(tourReview *model.TourReview) error {
	error := service.TourReviewRepository.UpdateReview(tourReview)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
