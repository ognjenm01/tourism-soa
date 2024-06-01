package handler

import (
	"context"
	"log"
	"strconv"
	"time"
	"tour/model"
	"tour/proto/tour"
	"tour/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TouristPositionHandler struct {
	tour.UnimplementedTouristPositionServiceServer
	TouristPositionService *service.TouristPositionService
}

func mapGrpcToModel(grpcPosition *tour.TouristPosition) *model.TouristPosition {
	modelPosition := model.TouristPosition{
		ID:           int(grpcPosition.Id),
		TourProgress: *MapGrpcProgressToModelProgress(grpcPosition.TourProgress),
		UserId:       int(grpcPosition.UserId),
		Latitude:     float64(grpcPosition.Latitude),
		Longitude:    float64(grpcPosition.Longitude),
		UpdatedAt:    time.Unix(grpcPosition.UpdatedAt.Seconds, int64(grpcPosition.UpdatedAt.Nanos)),
	}

	return &modelPosition
}

func (handler *TouristPositionHandler) CreateTouristPosition(ctx context.Context, request *tour.TouristPosition) (*tour.EmptyResponse, error) {
	result := mapGrpcToModel(request)
	err := handler.TouristPositionService.CreateTouristPosition(result)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TouristPositionHandler) GetTouristPositionById(ctx context.Context, request *tour.Id) (*tour.TouristPositionResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.TouristPositionService.GetTouristPositionById(id)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	grpcPosition := &tour.TouristPosition{
		Id:           int64(result.ID),
		TourProgress: mapModelProgressToGrpcProgress(&result.TourProgress),
		UserId:       int64(result.UserId),
		Latitude:     float32(result.Latitude),
		Longitude:    float32(result.Longitude),
		UpdatedAt:    &timestamppb.Timestamp{Seconds: result.UpdatedAt.Unix(), Nanos: int32(result.UpdatedAt.Nanosecond())},
	}

	return &tour.TouristPositionResponse{Position: grpcPosition}, nil
}

func (handler *TouristPositionHandler) GetTouristPositionByUser(ctx context.Context, request *tour.Id) (*tour.TouristPositionResponse, error) {
	result, err := handler.TouristPositionService.GetTouristPositionByUser(int(request.Id))

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	grpcPosition := &tour.TouristPosition{
		Id:           int64(result.ID),
		TourProgress: mapModelProgressToGrpcProgress(&result.TourProgress),
		UserId:       int64(result.UserId),
		Latitude:     float32(result.Latitude),
		Longitude:    float32(result.Longitude),
		UpdatedAt:    &timestamppb.Timestamp{Seconds: result.UpdatedAt.Unix(), Nanos: int32(result.UpdatedAt.Nanosecond())},
	}

	return &tour.TouristPositionResponse{Position: grpcPosition}, nil
}

func (handler *TouristPositionHandler) UpdateTouristPosition(ctx context.Context, request *tour.TouristPosition) (*tour.EmptyResponse, error) {
	result := mapGrpcToModel(request)
	err := handler.TouristPositionService.UpdateTouristPosition(result)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}
