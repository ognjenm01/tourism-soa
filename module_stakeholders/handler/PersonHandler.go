package handler

import (
	"context"
	"log"
	"stakeholder/model"
	stakeholders "stakeholder/proto"
	"stakeholder/service"
	"strconv"
)

type PersonHandler struct {
	stakeholders.UnimplementedPersonServiceServer
	PersonService *service.PersonService
}

func mapGrpcToModel(grpcPosition *stakeholders.Person) *model.Person {
	modelPosition := model.Person{
		ID:           int(grpcPosition.Id),
		UserId:       int(grpcPosition.UserId),
		Name:         string(grpcPosition.Name),
		Surname:      string(grpcPosition.Surname),
		Email:        string(grpcPosition.Email),
		ProfileImage: string(grpcPosition.ProfileImage),
		Biography:    string(grpcPosition.Biography),
		Quote:        string(grpcPosition.Quote),
	}
	return &modelPosition
}

func mapModelToGrpc(result *model.Person) *stakeholders.Person {
	grpcPosition := &stakeholders.Person{
		Id:           int64(result.ID),
		UserId:       int64(result.UserId),
		Name:         string(result.Name),
		Surname:      string(result.Surname),
		Email:        string(result.Email),
		ProfileImage: string(result.Email),
		Biography:    string(result.Biography),
		Quote:        string(result.Quote),
	}
	return grpcPosition
}

func (handler *PersonHandler) CreatePerson(ctx context.Context, request *stakeholders.Person) (*stakeholders.EmptyResponse, error) {
	result := mapGrpcToModel(request)
	err := handler.PersonService.CreatePerson(result)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &stakeholders.EmptyResponse{}, nil
}

func (handler *PersonHandler) GetPersonById(ctx context.Context, request *stakeholders.Id) (*stakeholders.PersonResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.PersonService.GetPersonById(id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	grpcPosition := mapModelToGrpc(result)
	return &stakeholders.PersonResponse{Person: grpcPosition}, nil
}

func (handler *PersonHandler) GetAllPeople(ctx context.Context, request *stakeholders.Empty) (*stakeholders.MultiPersonResponse, error) {
	result, err := handler.PersonService.GetAllPeople()
	var pMessage []*stakeholders.Person
	if err != nil {
		log.Print(err.Error())
		return &stakeholders.MultiPersonResponse{Persons: pMessage}, err
	}

	for _, p := range *result {
		grpcPosition := mapModelToGrpc(&p)
		pMessage = append(pMessage, grpcPosition)
	}

	return &stakeholders.MultiPersonResponse{Persons: pMessage}, nil
}
