package main

import (
	"log"
	"net/http"
	"tour/handler"
	"tour/infrastructure"
	"tour/repo"
	"tour/service"

	"github.com/gorilla/mux"
)

func startServer(handler *handler.TourHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/tours/{id}", handler.Find).Methods("GET")
	router.HandleFunc("/api/tours", handler.Create).Methods("POST")
	log.Println("[SERVER] - Server starting...")
	log.Fatal(http.ListenAndServe(":8080", router))
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
