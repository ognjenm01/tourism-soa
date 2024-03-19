package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/handler"
	"module_blog.xws.com/infrastructure"
	"module_blog.xws.com/repo"
	"module_blog.xws.com/service"
)

func startServer(blogHandler *handler.BlogHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")
	router.HandleFunc("/blogs", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogHandler.GetById).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func main() {
	database := infrastructure.InitDb()
	if database == nil {
		fmt.Println("1")
		log.Fatalln("Odjebi")
	}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)
}
