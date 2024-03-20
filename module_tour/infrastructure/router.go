package infrastructure

import (
	"tour/handler"

	"github.com/gorilla/mux"
)

func InitRouter(handler *handler.TourHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// TOURS
	router.HandleFunc("/api/tours/{id}", handler.GetTourById).Methods("GET")
	router.HandleFunc("/api/tours/byauthor/{id}", handler.GetTourByAuthor).Methods("GET")
	router.HandleFunc("/api/tours/bystatus", handler.GetToursByStatus).Methods("POST")
	router.HandleFunc("/api/tours", handler.CreateTour).Methods("POST")
	router.HandleFunc("/api/tours/{id}", handler.UpdateTour).Methods("PUT")
	router.HandleFunc("/api/tours", handler.GetAllTours).Methods("GET")

	// REVIEWS
	router.HandleFunc("/api/tourreview", handler.CreateReview).Methods("POST")
	router.HandleFunc("/api/tourreview", handler.GetAllReviews).Methods("GET")
	router.HandleFunc("/api/tourreview/{id}", handler.GetReviewById).Methods("GET")

	// KEYPOINTS
	router.HandleFunc("/api/keypoints", handler.CreateKeypoint).Methods("POST")
	router.HandleFunc("/api/keypoints/tour/{id}", handler.GetKeypointsByTourId).Methods("GET")
	router.HandleFunc("/api/keypoints", handler.CreateKeypoint).Methods("PUT")
	router.HandleFunc("/api/keypoints/{id}", handler.DeleteKeypoint).Methods("DELETE")

	// EQUIPMENT
	router.HandleFunc("/api/equipment", handler.CreateEquipment).Methods("POST")
	router.HandleFunc("/api/equipment", handler.GetAllEquipment).Methods("GET")
	router.HandleFunc("/api/equipment/tour/{id}", handler.GetEquipmentByTourId).Methods("GET")
	router.HandleFunc("/api/equipment/{id}", handler.UpdateEquipment).Methods("PUT")

	// TOUR EQUIPMENT
	router.HandleFunc("/api/tourequipment", handler.GetAllTourEquipment).Methods("GET")
	router.HandleFunc("/api/tourequipment", handler.CreateTourEquipment).Methods("POST")
	router.HandleFunc("/api/tourequipment/delete", handler.DeleteTourEquipment).Methods("POST")

	// TOUR PROGRESS
	router.HandleFunc("/api/tourprogress", handler.CreateTourProgress).Methods("POST")
	router.HandleFunc("/api/tourprogress", handler.GetAllTourProgress).Methods("GET")
	router.HandleFunc("/api/tourprogress/{id}", handler.GetTourProgressById).Methods("GET")
	router.HandleFunc("/api/tourprogress/{id}", handler.UpdateTourProgress).Methods("PUT")
	router.HandleFunc("/api/tourprogress/{id}", handler.DeleteTourProgress).Methods("DELETE")

	// TOURIST POSITION
	router.HandleFunc("/api/touristposition", handler.CreateTouristPosition).Methods("POST")
	router.HandleFunc("/api/touristposition/{id}", handler.GetTouristPositionById).Methods("GET")
	router.HandleFunc("/api/touristposition/byuser/{id}", handler.GetTouristPositionByUser).Methods("GET")
	router.HandleFunc("/api/touristposition/{id}", handler.UpdateTouristPosition).Methods("PUT")

	return router
}
