package service

import (
	"log"
	"stakeholder/model"
	"stakeholder/repo"
)

type UserService struct {
	UserRepository *repo.UserRepository
}

func (service *UserService) GetAllUsers() (*[]model.User, error) {
	users, error := service.UserRepository.GetAllUsers()
	if error != nil {
		log.Printf("[DB] - No tour reviews in db!\n")
		return nil, error
	}
	return &users, nil
}

func (service *UserService) GetUserById(id string) (*model.User, error) {
	user, error := service.UserRepository.GetUserById(id)
	if error != nil {
		log.Printf("[DB] - %s", error)
		return nil, error
	}

	return &user, nil
}

func (service *UserService) CreateUser(user *model.User) error {
	error := service.UserRepository.CreateUser(user)
	if error != nil {
		log.Printf("[DB] - nesto izuzetno pametno kod kreiranja korisnika")
		return error
	}
	return nil
}
