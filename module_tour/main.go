package main

import (
	"log"
	"net/http"
	"tour/handler"
	"tour/infrastructure"
	"tour/repo"
	"tour/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func startServer(handler *handler.TourHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/tours/{id}", handler.GetTourById).Methods("GET")
	router.HandleFunc("/api/tours/byauthor/{id}", handler.GetTourByAuthor).Methods("GET")
	router.HandleFunc("/api/tours/bystatus", handler.GetToursByStatus).Methods("POST")
	router.HandleFunc("/api/tours", handler.CreateTour).Methods("POST")
	router.HandleFunc("/api/tourreview", handler.CreateReview).Methods("POST")
	router.HandleFunc("/api/tourreview", handler.FindAllReviews).Methods("GET")
	router.HandleFunc("/api/tours/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/api/keypoints", handler.CreateKeypoint).Methods("POST")
	router.HandleFunc("/api/keypoints/tour/{id}", handler.GetKeypointsByTourId).Methods("GET")
	log.Println("[SERVER] - Server starting...")

	//FIXME: VERY UNSAFE! Add origins to env, fix the rest
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	hnd := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", hnd))
}

func main() {
	database := infrastructure.InitDB()
	if database == nil {
		log.Fatalln("[DB] - Failed to connect to database")
		return
	}
	log.Println("[DB] - Successfully connected to database")

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	keypointRepository := &repo.KeypointRepository{DatabaseConnection: database}

	tourService := &service.TourService{TourRepository: tourRepository}
	keypointService := &service.KeypointService{KeypointRepository: keypointRepository}

	tourReviewRepository := &repo.TourReviewRepository{DatabaseConnection: database}
	tourReviewService := &service.TourReviewService{TourReviewRepository: tourReviewRepository}

	handler := &handler.TourHandler{
		TourService:       tourService,
		KeypointService:   keypointService,
		TourReviewService: tourReviewService,
	}

	startServer(handler)
}
