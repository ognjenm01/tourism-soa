package handler

import (
	"context"
	"log"
	"strconv"
	"tour/model"
	"tour/proto/tour"
	"tour/service"
)

type KeypointHandler struct {
	tour.UnimplementedKeypointServiceServer
	KeypointService *service.KeypointService
}

func mapGrpcKeypointToModelKeypoint(grpcKeypoint *tour.Keypoint) *model.Keypoint {
	modelKeypoint := &model.Keypoint{
		ID:          int(grpcKeypoint.Id),
		TourID:      int(grpcKeypoint.TourId),
		Name:        grpcKeypoint.Name,
		Latitude:    float64(grpcKeypoint.Latitude),
		Longitude:   float64(grpcKeypoint.Longitude),
		Description: grpcKeypoint.Description,
		Position:    uint(grpcKeypoint.Position),
		Image:       grpcKeypoint.Image,
		Secret:      grpcKeypoint.Secret,
	}

	return modelKeypoint
}

func mapModelKeypointToGrpcKeypoint(k *model.Keypoint) *tour.Keypoint {
	grpcKeypoint := &tour.Keypoint{
		Id:          int64(k.ID),
		TourId:      int64(k.TourID),
		Name:        k.Name,
		Latitude:    float32(k.Latitude),
		Longitude:   float32(k.Longitude),
		Description: k.Description,
		Position:    uint64(k.Position),
		Image:       k.Image,
		Secret:      k.Secret,
	}

	return grpcKeypoint
}

func (handler *KeypointHandler) CreateKeypoint(ctx context.Context, request *tour.Keypoint) (*tour.EmptyResponse, error) {
	keypoint := mapGrpcKeypointToModelKeypoint(request)
	err := handler.KeypointService.CreateKeypoint(keypoint)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *KeypointHandler) GetKeypointsByTourId(ctx context.Context, request *tour.Id) (*tour.MultiKeypointResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	keypoints, err := handler.KeypointService.GetKeypointsByTourId(id)

	var protoMessages []*tour.Keypoint

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for _, k := range *keypoints {
		protoMessages = append(protoMessages, mapModelKeypointToGrpcKeypoint(&k))
	}

	return &tour.MultiKeypointResponse{Keypoint: protoMessages}, nil
}

func (handler *KeypointHandler) UpdateKeypoint(ctx context.Context, request *tour.Keypoint) (*tour.EmptyResponse, error) {
	modelKeypoint := mapGrpcKeypointToModelKeypoint(request)
	err := handler.KeypointService.CreateKeypoint(modelKeypoint)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *KeypointHandler) DeleteKeypoint(ctx context.Context, request *tour.Id) (*tour.EmptyResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	err := handler.KeypointService.DeleteKeypoint(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}
