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
	TourService          *service.TourService
	KeypointService      *service.KeypointService
	TourReviewService    *service.TourReviewService
	EquipmentService     *service.EquipmentService
	TourEquipmentService *service.TourEquipmentService
}

// ------------------------------------------------------------------------------------------- TOUR CRUD

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

//---------------------------------------------------------------------------------------- TOUR REVIEW CRUD

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

//---------------------------------------------------------------------------------------- EQUIPMENT CRUD

func (handler *TourHandler) CreateEquipment(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment

	error := json.NewDecoder(req.Body).Decode(&equipment)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.EquipmentService.CreateEquipment(&equipment)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) GetAllEquipment(writer http.ResponseWriter, req *http.Request) {
	equipment, error := handler.EquipmentService.GetAllEquipment()
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(equipment)
	}
}

func (handler *TourHandler) GetEquipmentById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	equipment, error := handler.EquipmentService.GetEquipmentById(id)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(equipment)
	}
}

func (handler *TourHandler) GetEquipmentByTourId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	equipment, error := handler.EquipmentService.GetEquipmentByTourId(id)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *TourHandler) UpdateEquipment(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	error := json.NewDecoder(req.Body).Decode(&equipment)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.EquipmentService.UpdateEquipment(&equipment)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

//---------------------------------------------------------------------------------------- TOUR EQUIPMENT CRUD

func (handler *TourHandler) CreateTourEquipment(writer http.ResponseWriter, req *http.Request) {
	var tourEquipment model.TourEquipment

	error := json.NewDecoder(req.Body).Decode(&tourEquipment)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.TourEquipmentService.CreateTourEquipment(&tourEquipment)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) GetAllTourEquipment(writer http.ResponseWriter, req *http.Request) {
	tourEquipment, error := handler.TourEquipmentService.GetAllTourEquipment()
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tourEquipment)
	}
}

func (handler *TourHandler) DeleteTourEquipment(writer http.ResponseWriter, req *http.Request) {
	var tourEquipment model.TourEquipment
	error := json.NewDecoder(req.Body).Decode(&tourEquipment)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	error = handler.TourEquipmentService.DeleteTourEquipment(&tourEquipment)
	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
