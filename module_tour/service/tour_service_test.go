package service

import (
	"testing"
	"time"
	"tour/model"
	"tour/repo"
	"tour/test_infrastructure"

	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db := test_infrastructure.InitDB()
	return db
}

func initTourService() TourService {
	return TourService{TourRepository: &repo.TourRepository{DatabaseConnection: initDB()}}
}

// Testing insert for tour
func TestCreateTour(t *testing.T) {
	tourService := initTourService()
	tour := model.Tour{ID: -1, UserId: -1, Name: "Debug tour", Description: "Debug test here", Price: 300.0, Difficulty: 1, TransportType: 1, Status: 1, Tags: []model.TourTag{}, Keypoints: []model.Keypoint{}, Duration: 5, Distance: 5, StatusUpdateTime: time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)}

	error := tourService.CreateTour(&tour)

	if error != nil {
		t.Fatal("Tour insert failed")
	}
}

// Testing getting all tours
func TestGetAll(t *testing.T) {
	tourService := initTourService()

	tours, error := tourService.GetAll()

	if error != nil {
		t.Fatalf("Getting all tours failed: %s", error)
	}

	if len(*tours) != 3 {
		t.Fatal("Wrong count of tours in db")
	}
}

// Testing getting tours by author
func TestGetToursByAuthor(t *testing.T) {
	tourService := initTourService()

	tours, error := tourService.GetToursByAuthor("-1")

	if error != nil {
		t.Fatalf("Getting tour by author id: %s", error)
	}

	if len(*tours) != 2 {
		t.Fatal("Wrong tour count")
	}

	if (*tours)[0].Name != "Tour to Spain" || (*tours)[1].Name != "Tour to Argentina" {
		t.Fatal("Wrong tours found")
	}
}

// Testing getting tour by statuses, example {PUBLISHED, DRAFT} as {0, 1}
func TestGetToursByStatuses(t *testing.T) {
	tourService := initTourService()

	tours, error := tourService.GetToursByStatus([]model.TourStatus{0})

	if error != nil {
		t.Fatalf("Error while getting tour by statuses %s", error)
	}

	if len(*tours) != 2 {
		t.Fatal("Wrong tour count")
	}
}

// Testing get tour by id
func TestGetTourById(t *testing.T) {
	tourService := initTourService()

	tour, error := tourService.GetTourById("1")

	if error != nil {
		t.Fatalf("Getting tour by id failed: %s", error)
	}

	if tour.Name != "Fantastic tour to Serbia" {
		t.Fatalf("Wrong name, expected 'Fantastic tour to Serbia', got %s", tour.Name)
	}

	if tour.Price != 300 {
		t.Fatalf("Wrong price, expected 300, got %f", tour.Price)
	}
}

func TestUpdateTour(t *testing.T) {
	tourService := initTourService()
	tour, error := tourService.GetTourById("1")

	if error != nil {
		t.Fatalf("Failed getting tour by id %s", error)
	}

	tour.Name = "Fantastic tour to Korea"

	tourService.UpdateTour(1, tour)

	tour, error = tourService.GetTourById("1")

	if error != nil {
		t.Fatalf("Failed getting tour by id %s", error)
	}

	if tour.Name != "Fantastic tour to Korea" {
		t.Fatalf("Expected tour name 'Fantastic tour to Korea', got %s", tour.Name)
	}

}
