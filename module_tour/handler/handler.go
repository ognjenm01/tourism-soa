package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tour/model"
	"tour/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService       *service.TourService
	KeypointService   *service.KeypointService
	TourReviewService *service.TourReviewService
}

// ------------------------------------------------------------------------------------------- TOUR CRUD

func (handler *TourHandler) GetTourById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tour, error := handler.TourService.GetTourById(id)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tour)
	}
}

func (handler *TourHandler) GetToursByStatus(writer http.ResponseWriter, req *http.Request) {
	var statuses []model.TourStatus
	error := json.NewDecoder(req.Body).Decode(&statuses)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tours, error := handler.TourService.GetToursByStatus(statuses)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil || len(*tours) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tours)
	}
}

func (handler *TourHandler) GetTourByAuthor(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tours, error := handler.TourService.GetToursByAuthor(id)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil || len(*tours) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tours)
	}
}

func (handler *TourHandler) CreateTour(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	error := json.NewDecoder(req.Body).Decode(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourService.CreateTour(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) UpdateTour(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	ID, error := strconv.Atoi(id)

	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var tour model.Tour
	error = json.NewDecoder(req.Body).Decode(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourService.UpdateTour(ID, &tour)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

//------------------------------------------------------------------------------------------- KEYPOINT CRUD

func (handler *TourHandler) GetKeypointsByTourId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	keypoints, error := handler.KeypointService.GetKeypointsByTourId(id)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keypoints)
}

func (handler *TourHandler) CreateKeypoint(writer http.ResponseWriter, req *http.Request) {
	var keypoint model.Keypoint
	error := json.NewDecoder(req.Body).Decode(&keypoint)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	error = handler.KeypointService.CreateKeypoint(&keypoint)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

//----------------------------------------------------------------------------------------TOUR REVIEW CRUD

func (handler *TourHandler) GetReviewById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tourReview, error := handler.TourReviewService.GetReviewById(id)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tourReview)
	}
}

func (handler *TourHandler) GetAllReviews(writer http.ResponseWriter, req *http.Request) {
	tourReviews, error := handler.TourReviewService.GetAllReviews()
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tourReviews)
	}
}

func (handler *TourHandler) CreateReview(writer http.ResponseWriter, req *http.Request) {
	var tourReview model.TourReview
	error := json.NewDecoder(req.Body).Decode(&tourReview)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourReviewService.CreateReview(&tourReview)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) UpdateReview(writer http.ResponseWriter, req *http.Request) {
	var tourReview model.TourReview
	error := json.NewDecoder(req.Body).Decode(&tourReview)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourReviewService.UpdateReview(&tourReview)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
