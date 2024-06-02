package handler

import (
	"context"
	"log"
	"stakeholder/model"
	stakeholders "stakeholder/proto"
	"stakeholder/service"
	"strconv"
)

type UserHandler struct {
	stakeholders.UnimplementedUserServiceServer
	UserService *service.UserService
}

func mapGrpcUserToModel(grpcPosition *stakeholders.User) *model.User {
	modelPosition := model.User{
		ID:        int(grpcPosition.Id),
		Username:  string(grpcPosition.Username),
		Password:  string(grpcPosition.Password),
		Role:      int(grpcPosition.Role),
		IsActive:  bool(grpcPosition.Isactive),
		IsBlocked: bool(grpcPosition.Isblocked),
		IsEnabled: bool(grpcPosition.Isenabled),
	}
	return &modelPosition
}

func mapModelUserToGrpc(result *model.User) *stakeholders.User {
	grpcPosition := &stakeholders.User{
		Id:        int64(result.ID),
		Username:  string(result.Username),
		Password:  string(result.Password),
		Role:      int64(result.Role),
		Isactive:  bool(result.IsActive),
		Isblocked: bool(result.IsBlocked),
		Isenabled: bool(result.IsEnabled),
	}
	return grpcPosition
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *stakeholders.User) (*stakeholders.EmptyResponse, error) {
	result := mapGrpcUserToModel(request)
	err := handler.UserService.CreateUser(result)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &stakeholders.EmptyResponse{}, nil
}

func (handler *UserHandler) GetUserById(ctx context.Context, request *stakeholders.Id) (*stakeholders.UserResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.UserService.GetUserById(id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	grpcPosition := mapModelUserToGrpc(result)
	return &stakeholders.UserResponse{User: grpcPosition}, nil
}

func (handler *UserHandler) GetUserByUsername(ctx context.Context, request *stakeholders.Username) (*stakeholders.UserResponse, error) {
	result, err := handler.UserService.GetUserByUsername(request.Username)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	grpcPosition := mapModelUserToGrpc(result)
	return &stakeholders.UserResponse{User: grpcPosition}, nil
}

func (handler *UserHandler) GetAllUsers(ctx context.Context, request *stakeholders.Empty) (*stakeholders.MultiUserResponse, error) {
	result, err := handler.UserService.GetAllUsers()
	var pMessage []*stakeholders.User
	if err != nil {
		log.Print(err.Error())
		return &stakeholders.MultiUserResponse{Users: pMessage}, err
	}

	for _, p := range *result {
		grpcPosition := mapModelUserToGrpc(&p)
		pMessage = append(pMessage, grpcPosition)
	}

	return &stakeholders.MultiUserResponse{Users: pMessage}, nil
}
