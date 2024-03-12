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
	TourReviewService *service.TourReviewService
}

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

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	error := json.NewDecoder(req.Body).Decode(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourService.Create(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
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

	error = handler.TourService.Update(ID, &tour)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

//*************************************************************

func (handler *TourHandler) FindReviewById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tourReview, error := handler.TourReviewService.FindReviewById(id)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tourReview)
	}
}

func (handler *TourHandler) FindAllReviews(writer http.ResponseWriter, req *http.Request) {
	tourReviews, error := handler.TourReviewService.FindAllReviews()
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
