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
	router.HandleFunc("/api/tours", handler.Create).Methods("POST")
	router.HandleFunc("/api/tours", handler.Update).Methods("PUT")
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

	repo := &repo.TourRepository{DatabaseConnection: database}
	service := &service.TourService{TourRepository: repo}
	handler := &handler.TourHandler{TourService: service}
	startServer(handler)
}
