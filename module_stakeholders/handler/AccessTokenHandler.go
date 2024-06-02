package handler

import (
	"context"
	"log"
	"stakeholder/model"
	stakeholders "stakeholder/proto"
	"stakeholder/service"
)

type AccessTokenHandler struct {
	stakeholders.UnimplementedAccessTokenServiceServer
	AccessTokenService *service.AccessTokenService
}

func mapGrpcAccessToModel(grpcPosition *stakeholders.AccessToken) *model.AccessToken {
	modelPosition := model.AccessToken{
		ID:          int(grpcPosition.Id),
		AccessToken: string(grpcPosition.Accesstoken),
	}
	return &modelPosition
}

func mapModelAccessToGrpc(result *model.AccessToken) *stakeholders.AccessToken {
	grpcPosition := &stakeholders.AccessToken{
		Id:          int64(result.ID),
		Accesstoken: string(result.AccessToken),
	}
	return grpcPosition
}

// mapGrpcLoginToModel maps the gRPC Login message to the model.Login struct
func mapGrpcLoginToModel(grpcLogin *stakeholders.Login) *model.Login {
	modelLogin := model.Login{
		Username: grpcLogin.Username,
		Password: grpcLogin.Password,
	}
	return &modelLogin
}

// mapModelLoginToGrpc maps the model.Login struct to the gRPC Login message
func mapModelLoginToGrpc(modelLogin *model.Login) *stakeholders.Login {
	grpcLogin := &stakeholders.Login{
		Username: modelLogin.Username,
		Password: modelLogin.Password,
	}
	return grpcLogin
}

// mapGrpcRegisterToModel maps the gRPC Register message to the model.Register struct
func mapGrpcRegisterToModel(grpcRegister *stakeholders.Register) *model.Register {
	modelRegister := model.Register{
		Username:     grpcRegister.Username,
		Password:     grpcRegister.Password,
		Name:         grpcRegister.Name,
		Surname:      grpcRegister.Surname,
		Email:        grpcRegister.Email,
		ProfileImage: grpcRegister.ProfileImage,
		Biography:    grpcRegister.Biography,
		Quote:        grpcRegister.Quote,
	}
	return &modelRegister
}

// mapModelRegisterToGrpc maps the model.Register struct to the gRPC Register message
func mapModelRegisterToGrpc(modelRegister *model.Register) *stakeholders.Register {
	grpcRegister := &stakeholders.Register{
		Username:     modelRegister.Username,
		Password:     modelRegister.Password,
		Name:         modelRegister.Name,
		Surname:      modelRegister.Surname,
		Email:        modelRegister.Email,
		ProfileImage: modelRegister.ProfileImage,
		Biography:    modelRegister.Biography,
		Quote:        modelRegister.Quote,
	}
	return grpcRegister
}

func (handler *AccessTokenHandler) LoginUser(ctx context.Context, request *stakeholders.Login) (*stakeholders.AccessToken, error) {
	result := mapGrpcLoginToModel(request)
	token, err := handler.AccessTokenService.LoginUser(result)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	if token == nil {
		return nil, nil
	}
	t := mapModelAccessToGrpc(token)
	return t, nil
}

func (handler *AccessTokenHandler) RegisterUser(ctx context.Context, request *stakeholders.Register) (*stakeholders.EmptyResponse, error) {
	result := mapGrpcRegisterToModel(request)
	err := handler.AccessTokenService.RegisterUser(result)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &stakeholders.EmptyResponse{}, nil
}
