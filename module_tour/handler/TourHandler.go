package handler

import (
	"encoding/json"
	"net/http"
	"tour/model"
	"tour/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
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
	var tour model.Tour
	error := json.NewDecoder(req.Body).Decode(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourService.Update(&tour)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
