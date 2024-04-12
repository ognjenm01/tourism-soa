package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/model"
	"module_blog.xws.com/service"
)

type KeyProduct struct{}

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	blog, err := handler.BlogService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (handler *BlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	blogs, error := handler.BlogService.GetAll()
	writer.Header().Set("Content-Type", "application/json")

	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(blogs)
	}
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {

	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.Create(&blog)
	if err != nil {
		println("Error while creating a blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) Update(writer http.ResponseWriter, req *http.Request) {
	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.Update(&blog)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	error := handler.BlogService.Delete(id)
	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *BlogHandler) AddRating(writer http.ResponseWriter, req *http.Request) {
	rating := model.BlogRating{}
	err := json.NewDecoder(req.Body).Decode(&rating)
	if err != nil {
		fmt.Println("lets go")
		//log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.AddRating(&rating)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
