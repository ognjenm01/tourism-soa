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

	// REVIEWS
	router.HandleFunc("/api/tourreview", handler.CreateReview).Methods("POST")
	router.HandleFunc("/api/tourreview", handler.GetAllReviews).Methods("GET")

	// KEYPOINTS
	router.HandleFunc("/api/keypoints", handler.CreateKeypoint).Methods("POST")
	router.HandleFunc("/api/keypoints/tour/{id}", handler.GetKeypointsByTourId).Methods("GET")

	return router
}
