package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/service"
)

type BlogRatingHandler struct {
	BlogRatingService *service.BlogRatingService
}

func (handler *BlogRatingHandler) GetByBlog(writer http.ResponseWriter, req *http.Request) {
	blogId := mux.Vars(req)["id"]

	ratings, err := handler.BlogRatingService.GetByBlog(blogId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(ratings)

}
