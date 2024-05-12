package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/model"
	"module_blog.xws.com/proto/blog"
	"module_blog.xws.com/service"
)

type KeyProduct struct{}

type BlogHandler struct {
	blog.UnimplementedBlogServiceServer
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

/*func convertSystemStatus(status model.SystemStatus) blog.Blog_SystemStatus {
	switch status {
	case model.DRAFT:
		return blog.Blog_DRAFT
	case model.PUBLISHED:
		return blog.Blog_PUBLISHED
	case model.CLOSED:
		return blog.Blog_CLOSED
	default:
		return blog.Blog_DRAFT
	}
}

func (handler *BlogHandler) GetById(ctx context.Context, request *blog.Id) (*blog.BlogResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.BlogService.GetById(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	protoResult := &blog.Blog{
		Kita:         int64(result.Kita),
		CreatorId:    int64(result.CreatorId),
		Title:        result.Title,
		Description:  result.Description,
		SystemStatus: convertSystemStatus(result.SystemStatus),
		ImageLinks:   result.ImageLinks,
		CreationDate: timestamppb.New(result.CreationDate),
		BlogStatuses: make([]*blog.BlogStatus, len(result.BlogStatuses)),
		BlogRatings:  make([]*blog.BlogRating, len(result.BlogRatings)),
	}

	for i, status := range result.BlogStatuses {
		protoBlogStatus := &blog.BlogStatus{
			Id:     int32(status.Id),
			BlogId: int32(status.BlogId),
			Name:   status.Name,
		}
		protoResult.BlogStatuses[i] = protoBlogStatus
	}

	for j, rating := range result.BlogRatings {
		protoBlogRating := &blog.BlogRating{
			Id:     int32(rating.Id),
			BlogId: int32(rating.BlogId),
			UserId: int32(rating.UserId),
			Rating: rating.Rating,
			//creationTime: rating.CreationTime,
		}
		protoResult.BlogRatings[j] = protoBlogRating
	}

	return &blog.BlogResponse{Blog: protoResult}, nil
}*/

/*func (handler *BlogHandler) GetAll(ctx context.Context, req *blog.Id) (*blog.MultiBlogResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	result, err := handler.BlogService.GetByPeopleUFollow(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

}*/

func (handler *BlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	blogs, error := handler.BlogService.GetByPeopleUFollow(id)
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
