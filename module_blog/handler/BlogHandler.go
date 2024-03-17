package handler

import (
	"encoding/json"
	"net/http"

	"module_blog.xws.com/model"
	"module_blog.xws.com/service"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing json")
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
