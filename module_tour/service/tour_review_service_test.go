package service

import (
	"testing"
	"tour/model"
	"tour/repo"
)

func initTourReviewService() TourReviewService {
	return TourReviewService{TourReviewRepository: &repo.TourReviewRepository{DatabaseConnection: initDB()}}
}

//TODO GET METODE KADA SKONTAS MOCK

// Testing create tour review
func TestCreateReviewEquipment(t *testing.T) {
	tourReviewService := initTourReviewService()
	tourReview := model.TourReview{}

	error := tourReviewService.CreateReview(&tourReview)

	if error != nil {
		t.Fatal("Failed review equipment")
	}
}

// TODO UPDATE
