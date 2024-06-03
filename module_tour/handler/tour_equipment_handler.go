package handler

import (
	"context"
	"log"
	"tour/model"
	"tour/proto/tour"
	"tour/service"
)

type TourEquipmentHandler struct {
	tour.UnimplementedTourEquipmentServiceServer
	TourEquipmentService *service.TourEquipmentService
}

func MapGrpcTourEquipmentModel(grpc *tour.TourEquipment) *model.TourEquipment {
	modelEq := &model.TourEquipment{
		ID:          int(grpc.Id),
		TourID:      int(grpc.Tourid),
		EquipmentID: int(grpc.Equipmentid),
	}

	return modelEq
}

func MapModelTourEquipmentGrpc(mdl *model.TourEquipment) *tour.TourEquipment {
	grpc := &tour.TourEquipment{
		Id:          int64(mdl.ID),
		Tourid:      int64(mdl.TourID),
		Equipmentid: int64(mdl.EquipmentID),
	}
	return grpc
}

func (handler *TourEquipmentHandler) CreateTourEquipment(ctx context.Context, req *tour.TourEquipment) (*tour.EmptyResponse, error) {
	tourEquipment := MapGrpcTourEquipmentModel(req)

	err := handler.TourEquipmentService.CreateTourEquipment(tourEquipment)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TourEquipmentHandler) GetAllTourEquipment(ctx context.Context, req *tour.Empty) (*tour.MultiTourEquipmentResponse, error) {
	tourEquipment, err := handler.TourEquipmentService.GetAllTourEquipment()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	tEqs := []*tour.TourEquipment{}
	for _, t := range *tourEquipment {
		tEqs = append(tEqs, MapModelTourEquipmentGrpc(&t))
	}
	return &tour.MultiTourEquipmentResponse{Tourequipment: tEqs}, nil
}

func (handler *TourEquipmentHandler) DeleteTourEquipment(ctx context.Context, request *tour.TourEquipment) (*tour.EmptyResponse, error) {
	tourEquipment := MapGrpcTourEquipmentModel(request)
	err := handler.TourEquipmentService.DeleteTourEquipment(tourEquipment)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &tour.EmptyResponse{}, nil
}
