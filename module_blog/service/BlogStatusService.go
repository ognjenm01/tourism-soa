package service

import (
	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogStatusService struct {
	BlogStatusRepo *repo.BlogStatusRepository
}

func (service *BlogStatusService) Create(status *model.BlogStatus) error {
	result := service.BlogStatusRepo.Create(status)

	if result != nil {
		return result
	}

	return nil
}
