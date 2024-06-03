package handler

import (
	"context"
	"log"
	"strconv"
	"tour/model"
	"tour/proto/tour"
	"tour/service"
)

type EquipmentHandler struct {
	tour.UnimplementedEquipmentServiceServer
	EquipmentService *service.EquipmentService
}

func MapGrpcEquipmentModel(grpc *tour.Equipment) *model.Equipment {
	modelEq := &model.Equipment{
		ID:          int(grpc.Id),
		Name:        grpc.Name,
		Description: grpc.Description,
	}

	return modelEq
}

func MapModelEquipmentToGrpc(modelEq *model.Equipment) *tour.Equipment {
	grpcEq := &tour.Equipment{
		Id:          int64(modelEq.ID),
		Name:        modelEq.Name,
		Description: modelEq.Description,
	}

	return grpcEq
}

func (handler *EquipmentHandler) CreateEquipment(ctx context.Context, req *tour.Equipment) (*tour.EmptyResponse, error) {
	equ := MapGrpcEquipmentModel(req)

	err := handler.EquipmentService.CreateEquipment(equ)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *EquipmentHandler) GetAllEquipment(ctx context.Context, req *tour.Empty) (*tour.MultiEquipmentResponse, error) {
	equ, err := handler.EquipmentService.GetAllEquipment()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	eqs := []*tour.Equipment{}
	for _, t := range *equ {
		eqs = append(eqs, MapModelEquipmentToGrpc(&t))
	}
	return &tour.MultiEquipmentResponse{Equipment: eqs}, nil
}

func (handler *EquipmentHandler) GetEquipmentById(ctx context.Context, req *tour.Id) (*tour.EquipmentResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	equ, err := handler.EquipmentService.GetEquipmentById(id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &tour.EquipmentResponse{Equipment: MapModelEquipmentToGrpc(equ)}, nil
}

func (handler *EquipmentHandler) GetEquipmentByTourId(ctx context.Context, req *tour.Id) (*tour.MultiEquipmentResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	equ, err := handler.EquipmentService.GetEquipmentByTourId(id)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	eqs := []*tour.Equipment{}
	for _, t := range *equ {
		eqs = append(eqs, MapModelEquipmentToGrpc(&t))
	}

	return &tour.MultiEquipmentResponse{Equipment: eqs}, nil
}

func (handler *EquipmentHandler) UpdateEquipment(ctx context.Context, req *tour.Equipment) (*tour.EmptyResponse, error) {
	equ := MapGrpcEquipmentModel(req)

	err := handler.EquipmentService.UpdateEquipment(equ)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}
