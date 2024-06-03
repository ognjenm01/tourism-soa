package handler

import (
	"context"
	"log"
	"strconv"
	"tour/model"
	"tour/proto/tour"
	"tour/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TourReviewHandler struct {
	tour.UnimplementedTourReviewServiceServer
	TourReviewService *service.TourReviewService
}

func mapGrpcReviewModel(grpc *tour.TourReview) *model.TourReview {
	tr := &model.TourReview{
		ID:         int(grpc.Id),
		Rating:     int(grpc.Rating),
		Comment:    grpc.Comment,
		VisitDate:  grpc.Visitdate.AsTime(),
		RatingDate: grpc.Ratingdate.AsTime(),
		ImageLinks: make([]model.ImageLink, len(grpc.Imagelinks)),
		UserId:     int(grpc.Userid),
		TourId:     int(grpc.Tourid),
	}

	uinfo := &model.UserInfo{
		FirstName: grpc.Userinfo.Firstname,
		LastName:  grpc.Userinfo.Lastname,
	}

	tr.UserInfo = *uinfo

	for i, im := range grpc.Imagelinks {
		tr.ImageLinks[i] = model.ImageLink{
			ID:           int(im.Id),
			TourReviewID: int(im.Tourreviewid),
			Link:         im.Link,
		}
	}

	return tr
}

func mapModelReviewToGrpc(rev *model.TourReview) *tour.TourReview {
	grpcReview := &tour.TourReview{
		Id:         int64(rev.ID),
		Rating:     int64(rev.Rating),
		Comment:    rev.Comment,
		Visitdate:  timestamppb.New(rev.VisitDate),
		Ratingdate: timestamppb.New(rev.RatingDate),
		Userid:     int64(rev.UserId),
		Tourid:     int64(rev.TourId),
		Imagelinks: make([]*tour.ImageLink, len(rev.ImageLinks)),
		Userinfo: &tour.UserInfo{
			Firstname: rev.UserInfo.FirstName,
			Lastname:  rev.UserInfo.LastName,
		},
	}

	for i, im := range rev.ImageLinks {
		grpcReview.Imagelinks[i] = &tour.ImageLink{
			Id:           int64(im.ID),
			Tourreviewid: int64(im.TourReviewID),
			Link:         im.Link,
		}
	}

	return grpcReview
}

func (handler *TourReviewHandler) CreateReview(ctx context.Context, req *tour.TourReview) (*tour.EmptyResponse, error) {
	tourReview := mapGrpcReviewModel(req)

	err := handler.TourReviewService.CreateReview(tourReview)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}

func (handler *TourReviewHandler) GetAllReviews(ctx context.Context, req *tour.Empty) (*tour.MultiTourReviewResponse, error) {
	tourReviews, err := handler.TourReviewService.GetAllReviews()
	//FIXME
	/*for i, review := range *tourReviews {


	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, error := http.Get(fmt.Sprintf("http://explorer:80/api/profile/userinfo/%d", review.UserId))
		if error != nil {
			writer.WriteHeader(http.StatusFailedDependency)
			return
		}
		defer resp.Body.Close()
		error = json.NewDecoder(resp.Body).Decode(&(*tourReviews)[i].UserInfo)
		if error != nil {
			writer.WriteHeader(http.StatusFailedDependency)
			return
		}
	}*/

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	trs := []*tour.TourReview{}
	for _, t := range *tourReviews {
		trs = append(trs, mapModelReviewToGrpc(&t))
	}
	return &tour.MultiTourReviewResponse{Tourreviews: trs}, nil
}

func (handler *TourReviewHandler) GetReviewById(ctx context.Context, req *tour.Id) (*tour.TourReviewResponse, error) {
	id := strconv.FormatInt(req.Id, 10)
	tourReview, err := handler.TourReviewService.GetReviewById(id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	//Fixme
	/*resp, error := http.Get(fmt.Sprintf("https://localhost:44333/api/profile/userinfo/%d", tourReview.ID))
	if error != nil {
		writer.WriteHeader(http.StatusFailedDependency)
		return
	}
	defer resp.Body.Close()
	error = json.NewDecoder(resp.Body).Decode(&tourReview.UserInfo)

	writer.Header().Set("Content-Type", "application/json")
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tourReview)
	}*/
	t := mapModelReviewToGrpc(tourReview)
	return &tour.TourReviewResponse{Tourreview: t}, nil
}

func (handler *TourReviewHandler) UpdateReview(ctx context.Context, req *tour.TourReview) (*tour.EmptyResponse, error) {
	tourReview := mapGrpcReviewModel(req)

	err := handler.TourReviewService.UpdateReview(tourReview)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &tour.EmptyResponse{}, nil
}
