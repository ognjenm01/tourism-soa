package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"tour/handler"
	"tour/infrastructure"
	"tour/proto/tour"
	"tour/repo"
	"tour/service"

	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	port := os.Getenv("TOURS_APP_PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	log.Printf("[SERVER] - Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, hnd))
}

func main() {
	database := infrastructure.InitDB()
	if database == nil {
		log.Fatalln("[DB] - Failed to connect to database")
		return
	}

	log.Println("[DB] - Successfully connected to database")

	/*tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}

	tourReviewRepository := &repo.TourReviewRepository{DatabaseConnection: database}
	tourReviewService := &service.TourReviewService{TourReviewRepository: tourReviewRepository}

	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepository}

	tourEquipmentRepository := &repo.TourEquipmentRepository{DatabaseConnection: database}
	tourEquipmentService := &service.TourEquipmentService{TourEquipmentRepository: tourEquipmentRepository}


	handler := &handler.TourHandler{
		TourService:            tourService,
		KeypointService:        keypointService,
		TourReviewService:      tourReviewService,
		EquipmentService:       equipmentService,
		TourEquipmentService:   tourEquipmentService,
		TourProgressService:    tourProgressService,
		TouristPositionService: touristPositionService,
	}*/

	tourProgressRepository := &repo.TourProgressRepository{DatabaseConnection: database}
	tourProgressService := &service.TourProgressService{TourProgressRepository: tourProgressRepository}
	tourProgressHandler := &handler.TourProgressHandler{TourProgressService: tourProgressService}

	touristPositionRepository := &repo.TouristPositionRepository{DatabaseConnection: database}
	touristPositionService := &service.TouristPositionService{TouristPositionRepository: touristPositionRepository}
	touristPositionHandler := &handler.TouristPositionHandler{TouristPositionService: touristPositionService}

	keypointRepository := &repo.KeypointRepository{DatabaseConnection: database}
	keypointService := &service.KeypointService{KeypointRepository: keypointRepository}
	keypointHandler := &handler.KeypointHandler{KeypointService: keypointService}

	port := ":" + os.Getenv("TOURS_APP_PORT")
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	fmt.Println("Registered gRPC server")

	fmt.Println("Listening at:" + port)

	tour.RegisterTourProgressServiceServer(grpcServer, tourProgressHandler)
	tour.RegisterTouristPositionServiceServer(grpcServer, touristPositionHandler)
	tour.RegisterKeypointServiceServer(grpcServer, keypointHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("Serving gRPC")
	grpcServer.Serve(lis)

	select {} //nzm
	//startServer(handler)
}
