package service

import (
	"fmt"
	"log"

	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}

func (service *BlogService) GetById(id string) (*model.Blog, error) {
	blog, err := service.BlogRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &blog, err

}

func (service *BlogService) GetAll() (*[]model.Blog, error) {
	blogs, error := service.BlogRepo.GetAll()

	if error != nil {
		log.Fatalf("Database is empty\n")
		return &[]model.Blog{}, error
	}

	return &blogs, nil
}

func (service *BlogService) Create(blog *model.Blog) error {
	result := service.BlogRepo.Create(blog)
	if result != nil {
		return result
	}

	return nil
}

func (service *BlogService) Update(blog *model.Blog) error {
	result := service.BlogRepo.Update(blog)
	if result != nil {
		return result
	}

	return nil
}

func (service *BlogService) Delete(id string) error {
	result := service.BlogRepo.Delete(id)
	if result != nil {
		return result
	}

	return nil
}
