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

type NewTourHandler struct {
	//fixme
	tour.UnimplementedTourProgressServiceServer
	TourService *service.TourService
}

func MapGrpcTourToModel(grpcProgress *tour.Tour) *model.Tour {
	tm := &model.Tour{
		ID:               int(grpcProgress.Id),
		UserId:           int(grpcProgress.Userid),
		Name:             string(grpcProgress.Name),
		Description:      string(grpcProgress.Description),
		Price:            float64(grpcProgress.Price),
		Difficulty:       int(grpcProgress.Difficulty),
		TransportType:    int(grpcProgress.Transporttype),
		Keypoints:        make([]model.Keypoint, len(grpcProgress.Keypoints)),
		Duration:         int(grpcProgress.Duration),
		Distance:         float64(grpcProgress.Distance),
		StatusUpdateTime: time.Unix(grpcProgress.Statusupdatetime.Seconds, int64(grpcProgress.Statusupdatetime.Nanos)),
		TourProgress:     make([]model.TourProgress, len(grpcProgress.Tourprogress)),
	}

	for i, kp := range grpcProgress.Keypoints {
		tm.Keypoints[i] = model.Keypoint{
			ID:          int(kp.Id),
			TourID:      int(kp.TourId),
			Name:        kp.Name,
			Latitude:    float64(kp.Latitude),
			Longitude:   float64(kp.Longitude),
			Description: kp.Description,
			Position:    uint(kp.Position),
			Image:       kp.Image,
			Secret:      kp.Secret,
		}
	}

	for _, tp := range grpcProgress.Tourprogress {
		startTime := time.Unix(tp.StartTime.Seconds, int64(tp.StartTime.Nanos))
		lastActivity := time.Unix(tp.LastActivity.Seconds, int64(tp.LastActivity.Nanos))
		tourProgress := model.TourProgress{
			ID:                int(tp.Id),
			TouristPositionId: int(tp.TouristPositionId),
			TourId:            int(tp.TourId),
			Status:            int(tp.Status),
			StartTime:         startTime,
			LastActivity:      lastActivity,
			CurrentKeypoint:   int(tp.CurrentKeypoint),
			IsInProgress:      tp.IsInProgress,
		}
		tm.TourProgress = append(tm.TourProgress, tourProgress)
	}

	return tm
}

func mapModelTourToGrpc(modelTour *model.Tour) *tour.Tour {
	grpcTour := &tour.Tour{
		Id:            int64(modelTour.ID),
		Userid:        int64(modelTour.UserId),
		Name:          modelTour.Name,
		Description:   modelTour.Description,
		Price:         float32(modelTour.Price),
		Difficulty:    int64(modelTour.Difficulty),
		Transporttype: int64(modelTour.TransportType),
		Keypoints:     make([]*tour.Keypoint, len(modelTour.Keypoints)),
		Tourprogress:  make([]*tour.TourProgress, len(modelTour.TourProgress)),
	}

	for i, kp := range modelTour.Keypoints {
		grpcTour.Keypoints[i] = &tour.Keypoint{
			Id:          int64(kp.ID),
			TourId:      int64(kp.TourID),
			Name:        kp.Name,
			Latitude:    float32(kp.Latitude),
			Longitude:   float32(kp.Longitude),
			Description: kp.Description,
			Position:    uint64(kp.Position),
			Image:       kp.Image,
			Secret:      kp.Secret,
		}
	}

	for i, tp := range modelTour.TourProgress {
		grpcTour.Tourprogress[i] = &tour.TourProgress{
			Id:                int64(tp.ID),
			TouristPositionId: int64(tp.TouristPositionId),
			TourId:            int64(tp.TourId),
			Status:            int64(tp.Status),
			StartTime:         &timestamppb.Timestamp{Seconds: tp.StartTime.Unix(), Nanos: int32(tp.StartTime.Nanosecond())},
			LastActivity:      &timestamppb.Timestamp{Seconds: tp.LastActivity.Unix(), Nanos: int32(tp.LastActivity.Nanosecond())},
			CurrentKeypoint:   int64(tp.CurrentKeypoint),
			IsInProgress:      tp.IsInProgress,
		}
	}

	return grpcTour
}

func (handler *TourHandler) CreateTour(ctx context.Context, to *tour.Tour) (*tour.EmptyResponse, error) {
	t := MapGrpcTourToModel(to)
	err := handler.TourService.CreateTour(t)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &tour.EmptyResponse{}, nil
}

func (handler *TourHandler) GetTourById(ctx context.Context, req *tour.Id) (*tour.TourResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	tourOld, err := handler.TourService.GetTourById(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	tt := mapModelTourToGrpc(tourOld)
	return &tour.TourResponse{Tour: tt}, nil
}

func (handler *TourHandler) GetToursByStatus(ctx context.Context, req []*tour.Status) (*tour.MultiTourResponse, error) {
	var statuses []int
	for _, st := range req {
		statuses = append(statuses, int(st.Status))
	}

	tours, err := handler.TourService.GetToursByStatus(statuses)
	if err != nil {
		return nil, err
	}

	finalTt := []*tour.Tour{}
	for _, t := range *tours {
		finalTt = append(finalTt, mapModelTourToGrpc(&t))
	}
	return &tour.MultiTourResponse{Tours: finalTt}, nil
}

func (handler *TourHandler) GetTourByAuthor(ctx context.Context, req *tour.Id) (*tour.MultiTourResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	tours, err := handler.TourService.GetToursByAuthor(id)
	if err != nil {
		return nil, err
	}
	tt := []*tour.Tour{}
	for _, t := range *tours {
		tt = append(tt, mapModelTourToGrpc(&t))
	}
	return &tour.MultiTourResponse{Tours: tt}, nil
}

func (handler *TourHandler) UpdateTour(ctx context.Context, req *tour.Tour) (*tour.EmptyResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	ID, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	to := MapGrpcTourToModel(req)

	err = handler.TourService.UpdateTour(ID, to)
	if err != nil {
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TourHandler) GetAllTours(ctx context.Context, req *tour.EmptyResponse) (*tour.MultiTourResponse, error) {
	tours, err := handler.TourService.GetAll()
	if err != nil {
		return nil, err
	}
	tt := []*tour.Tour{}
	for _, t := range *tours {
		tt = append(tt, mapModelTourToGrpc(&t))
	}
	return &tour.MultiTourResponse{Tours: tt}, nil
}
