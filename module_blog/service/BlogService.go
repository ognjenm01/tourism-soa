package service

import (
	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}

func (service *BlogService) Create(blog *model.Blog) error {
	result := service.BlogRepo.Create(blog)
	if result != nil {
		return result
	}

	return nil
}
