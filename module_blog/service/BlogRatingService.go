package service

import (
	"log"

	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogRatingService struct {
	BlogRatingRepo *repo.BlogRatingRepository
}

func (service *BlogRatingService) GetByBlog(blogId string) (*[]model.BlogRating, error) {
	ratings, error := service.BlogRatingRepo.GetByBlog(blogId)

	if error != nil {
		log.Fatalf("No rows found")
		return &[]model.BlogRating{}, error
	}

	return &ratings, nil
}
