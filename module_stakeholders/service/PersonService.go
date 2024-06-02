package service

import (
	"log"
	"stakeholder/model"
	"stakeholder/repo"
)

type PersonService struct {
	PersonRepository *repo.PersonRepository
}

func (service *PersonService) GetAllPeople() (*[]model.Person, error) {
	persons, error := service.PersonRepository.GetAllPeople()
	if error != nil {
		log.Printf("[DB] - No tour reviews in db!\n")
		return nil, error
	}
	return &persons, nil
}

func (service *PersonService) GetPersonById(id string) (*model.Person, error) {
	person, error := service.PersonRepository.GetPersonById(id)
	if error != nil {
		log.Printf("[DB] - %s", error)
		return nil, error
	}

	return &person, nil
}

func (service *PersonService) CreatePerson(person *model.Person) error {
	error := service.PersonRepository.CreatePerson(person)
	if error != nil {
		log.Printf("[DB] - nesto izuzetno pametno kod kreiranja osobe")
		return error
	}
	return nil
}
