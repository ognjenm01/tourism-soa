package service

import (
	"testing"
	"time"
	"tour/model"
	"tour/repo"
	"tour/test_infrastructure"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db := test_infrastructure.InitDB()
	test_infrastructure.ClearDB(db)
	return db
}

// Testing insert for tour
func TestCreateTour(t *testing.T) {
	tourService := TourService{TourRepository: &repo.TourRepository{DatabaseConnection: InitDB()}}
	tour := model.Tour{ID: -1, UserId: -1, Name: "Debug tour", Description: "Debug test here", Price: 300.0, Difficulty: 1, TransportType: 1, Status: 1, Tags: []model.TourTag{}, Keypoints: []model.Keypoint{}, Duration: 5, Distance: 5, StatusUpdateTime: time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)}

	error := tourService.CreateTour(&tour)
	if error != nil {
		t.Fatal("Tour insert failed")
	}
}
