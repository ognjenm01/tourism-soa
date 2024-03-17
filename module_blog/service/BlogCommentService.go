package service

import (
	"log"

	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogCommentService struct {
	BlogCommentRepository *repo.BlogCommentRepository
}

func (service *BlogCommentService) GetCommentById(id string) (*model.BlogComment, error) {
	blogComment, error := service.BlogCommentRepository.GetById(id)

	if error != nil {
		log.Fatalf("[DB] - No blog comment in db!\n")
		return nil, error
	}

	return &blogComment, nil
}

func (service *BlogCommentService) GetAllComments() (*[]model.BlogComment, error) {
	blogComments, error := service.BlogCommentRepository.GetAll()

	if error != nil {
		log.Fatalf("[DB] - No blog comments in db!\n")
		return nil, error
	}

	return &blogComments, nil
}

func (service *BlogCommentService) CreateComment(blogComment *model.BlogComment) error {
	error := service.BlogCommentRepository.Create(blogComment)

	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}

	return nil
}

func (service *BlogCommentService) UpdateComment(blogComment *model.BlogComment) error {
	error := service.BlogCommentRepository.Update(blogComment)

	if error != nil {
		log.Fatalf("[DB] - %s", error)
		return error
	}

	return nil
}
