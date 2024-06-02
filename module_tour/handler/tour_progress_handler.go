package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"tour/model"
	"tour/proto/tour"
	"tour/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TourProgressHandler struct {
	tour.UnimplementedTourProgressServiceServer
	TourProgressService *service.TourProgressService
}

func MapGrpcProgressToModelProgress(grpcProgress *tour.TourProgress) *model.TourProgress {
	modelProgress := &model.TourProgress{
		ID:                int(grpcProgress.Id),
		TouristPositionId: int(grpcProgress.TouristPositionId),
		TourId:            int(grpcProgress.TourId),
		Status:            int(grpcProgress.Status),
		StartTime:         time.Unix(grpcProgress.StartTime.Seconds, int64(grpcProgress.StartTime.Nanos)),
		LastActivity:      time.Unix(grpcProgress.LastActivity.Seconds, int64(grpcProgress.LastActivity.Nanos)),
		CurrentKeypoint:   int(grpcProgress.CurrentKeypoint),
		IsInProgress:      bool(grpcProgress.IsInProgress),
	}

	return modelProgress
}

func mapModelProgressToGrpcProgress(modelProgress *model.TourProgress) *tour.TourProgress {
	grpcProgress := &tour.TourProgress{
		Id:                int64(modelProgress.ID),
		TouristPositionId: int64(modelProgress.TouristPositionId),
		TourId:            int64(modelProgress.TourId),
		Status:            int64(modelProgress.Status),
		StartTime:         &timestamppb.Timestamp{Seconds: modelProgress.StartTime.Unix(), Nanos: int32(modelProgress.StartTime.Nanosecond())},
		LastActivity:      &timestamppb.Timestamp{Seconds: modelProgress.LastActivity.Unix(), Nanos: int32(modelProgress.LastActivity.Nanosecond())},
		CurrentKeypoint:   int64(modelProgress.CurrentKeypoint),
		IsInProgress:      modelProgress.IsInProgress,
	}

	return grpcProgress
}

func (handler *TourProgressHandler) CreateTourProgress(ctx context.Context, request *tour.TourProgress) (*tour.EmptyResponse, error) {
	newProgress := MapGrpcProgressToModelProgress(request)
	err := handler.TourProgressService.CreateTourProgress(newProgress)
	fmt.Println("usao")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TourProgressHandler) GetAllTourProgress(ctx context.Context, request *tour.Empty) (*tour.MultiTourProgressResponse, error) {
	result, err := handler.TourProgressService.GetAllTourProgress()

	var protoMessages []*tour.TourProgress

	if err != nil {
		log.Fatalln(err)
		return &tour.MultiTourProgressResponse{TourProgress: protoMessages}, err
	}

	for _, t := range *result {
		grpcProgress := &tour.TourProgress{
			Id:                int64(t.ID),
			TouristPositionId: int64(t.TouristPositionId),
			TourId:            int64(t.TourId),
			Status:            int64(t.Status),
			StartTime:         &timestamppb.Timestamp{Seconds: t.StartTime.Unix(), Nanos: int32(t.StartTime.Nanosecond())},
			LastActivity:      &timestamppb.Timestamp{Seconds: t.LastActivity.Unix(), Nanos: int32(t.LastActivity.Nanosecond())},
			CurrentKeypoint:   int64(t.CurrentKeypoint),
			IsInProgress:      t.IsInProgress,
		}

		protoMessages = append(protoMessages, grpcProgress)
	}

	return &tour.MultiTourProgressResponse{TourProgress: protoMessages}, nil
}

func (handler *TourProgressHandler) GetTourProgressById(ctx context.Context, request *tour.Id) (*tour.TourProgressResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.TourProgressService.GetTourProgressById(id)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	grpcProgress := &tour.TourProgress{
		Id:                int64(result.ID),
		TouristPositionId: int64(result.TouristPositionId),
		TourId:            int64(result.TourId),
		Status:            int64(result.Status),
		StartTime:         &timestamppb.Timestamp{Seconds: result.StartTime.Unix(), Nanos: int32(result.StartTime.Nanosecond())},
		LastActivity:      &timestamppb.Timestamp{Seconds: result.LastActivity.Unix(), Nanos: int32(result.LastActivity.Nanosecond())},
		CurrentKeypoint:   int64(result.CurrentKeypoint),
		IsInProgress:      result.IsInProgress,
	}

	return &tour.TourProgressResponse{TourProgress: grpcProgress}, nil
}

func (handler *TourProgressHandler) UpdateTourProgress(ctx context.Context, request *tour.TourProgress) (*tour.EmptyResponse, error) {
	progress := MapGrpcProgressToModelProgress(request)
	err := handler.TourProgressService.UpdateTourProgress(progress)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TourProgressHandler) DeleteTourProgress(ctx context.Context, request *tour.Id) (*tour.EmptyResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	err := handler.TourProgressService.DeleteTourProgress(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}
