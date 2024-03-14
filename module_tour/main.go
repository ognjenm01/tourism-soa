package main

import (
	"log"
	"net/http"
	"tour/handler"
	"tour/infrastructure"
	"tour/repo"
	"tour/service"

	"github.com/rs/cors"
)

func startServer(handler *handler.TourHandler) {
	log.Println("[SERVER] - Server starting...")

	//FIXME: VERY UNSAFE! Add origins to env, fix the rest
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	hnd := c.Handler(infrastructure.InitRouter(handler))
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
	tourService := &service.TourService{TourRepository: tourRepository}

	keypointRepository := &repo.KeypointRepository{DatabaseConnection: database}
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
