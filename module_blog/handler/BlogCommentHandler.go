package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/model"
	"module_blog.xws.com/service"
)

type BlogCommentHandler struct {
	BlogCommentService *service.BlogCommentService
}

func (handler *BlogCommentHandler) CreateComment(writer http.ResponseWriter, req *http.Request) {
	var blogComment model.BlogComment
	error := json.NewDecoder(req.Body).Decode(&blogComment)

	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.BlogCommentService.CreateComment(&blogComment)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-type", "application/json")
}

func (handler *BlogCommentHandler) GetAllComments(writer http.ResponseWriter, req *http.Request) {
	blogComments, error := handler.BlogCommentService.GetAllComments()
	writer.Header().Set("Content-type", "application/json")

	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(blogComments)
	}
}

func (handler *BlogCommentHandler) GetCommentById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	blogComment, error := handler.BlogCommentService.GetCommentById(id)
	writer.Header().Set("Content-type", "application/json")

	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(blogComment)
	}
}

func (handler *BlogCommentHandler) UpdateComment(writer http.ResponseWriter, req *http.Request) {
	var blogComment model.BlogComment
	error := json.NewDecoder(req.Body).Decode(&blogComment)

	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	error = handler.BlogCommentService.UpdateComment(&blogComment)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-type", "application/json")
}

func (handler *BlogCommentHandler) DeleteComment(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	error := handler.BlogCommentService.DeleteComment(id)

	if error != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
