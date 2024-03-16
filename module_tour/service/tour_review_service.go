package service

import (
	"log"
	"tour/model"
	"tour/repo"
)

type TourReviewService struct {
	TourReviewRepository *repo.TourReviewRepository
}

func (service *TourReviewService) GetReviewById(id string) (*model.TourReview, error) {
	tourReview, error := service.TourReviewRepository.GetById(id)
	if error != nil {
		//log.Fatalf("[DB] - No tour review in db!\n")
		return nil, error
	}
	return &tourReview, nil
}

func (service *TourReviewService) GetAllReviews() (*[]model.TourReview, error) {
	tourReviews, error := service.TourReviewRepository.GetAll()
	if error != nil {
		//log.Fatalf("[DB] - No tour reviews in db!\n")
		return nil, error
	}
	return &tourReviews, nil
}

func (service *TourReviewService) CreateReview(tourReview *model.TourReview) error {
	error := service.TourReviewRepository.Create(tourReview)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}

func (service *TourReviewService) UpdateReview(tourReview *model.TourReview) error {
	error := service.TourReviewRepository.Update(tourReview)
	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}
	return nil
}
