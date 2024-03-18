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
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
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

	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepository}

	tourEquipmentRepository := &repo.TourEquipmentRepository{DatabaseConnection: database}
	tourEquipmentService := &service.TourEquipmentService{TourEquipmentRepository: tourEquipmentRepository}

	tourProgressRepository := &repo.TourProgressRepository{DatabaseConnection: database}
	tourProgressService := &service.TourProgressService{TourProgressRepository: tourProgressRepository}

	touristPositionRepository := &repo.TouristPositionRepository{DatabaseConnection: database}
	touristPositionService := &service.TouristPositionService{TouristPositionRepository: touristPositionRepository}

	handler := &handler.TourHandler{
		TourService:            tourService,
		KeypointService:        keypointService,
		TourReviewService:      tourReviewService,
		EquipmentService:       equipmentService,
		TourEquipmentService:   tourEquipmentService,
		TourProgressService:    tourProgressService,
		TouristPositionService: touristPositionService,
	}

	startServer(handler)
}
