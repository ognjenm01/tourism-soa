package handler

import (
	"context"
	"log"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"module_blog.xws.com/model"
	"module_blog.xws.com/proto/blog"
	"module_blog.xws.com/service"
)

type BlogHandler struct {
	blog.UnimplementedBlogServiceServer
	BlogService *service.BlogService
}

func convertSystemStatus(status model.SystemStatus) blog.Blog_SystemStatus {
	switch status {
	case model.DRAFT:
		return blog.Blog_DRAFT
	case model.PUBLISHED:
		return blog.Blog_PUBLISHED
	case model.CLOSED:
		return blog.Blog_CLOSED
	default:
		return blog.Blog_DRAFT
	}
}

func (handler *BlogHandler) GetById(ctx context.Context, request *blog.Id) (*blog.BlogResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	result, err := handler.BlogService.GetById(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	protoResult := &blog.Blog{
		Kita:         int64(result.Kita),
		CreatorId:    int64(result.CreatorId),
		Title:        result.Title,
		Description:  result.Description,
		SystemStatus: convertSystemStatus(result.SystemStatus),
		ImageLinks:   result.ImageLinks,
		CreationDate: timestamppb.New(result.CreationDate),
		//BlogStatuses: make([]*blog.BlogStatus, len(result.BlogStatuses)),
		//BlogRatings:  make([]*blog.BlogRating, len(result.BlogRatings)),
	}

	/*for i, status := range result.BlogStatuses {
		protoBlogStatus := &blog.BlogStatus{
			Id:     int32(status.Id),
			BlogId: int32(status.BlogId),
			Name:   status.Name,
		}
		protoResult.BlogStatuses[i] = protoBlogStatus
	}

	for j, rating := range result.BlogRatings {
		protoBlogRating := &blog.BlogRating{
			Id:           int32(rating.Id),
			BlogId:       int32(rating.BlogId),
			UserId:       int32(rating.UserId),
			Rating:       rating.Rating,
			CreationTime: timestamppb.New(rating.CreationTime),
		}
		protoResult.BlogRatings[j] = protoBlogRating
	}*/

	return &blog.BlogResponse{Blog: protoResult}, nil
}

func (handler *BlogHandler) GetAll(ctx context.Context, req *blog.Empty) (*blog.MultiBlogResponse, error) {
	//id := strconv.FormatInt(req.Id, 10)
	result, err := handler.BlogService.GetByPeopleUFollow()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var protoMessages []*blog.Blog

	for _, b := range *result {
		protoResult := &blog.Blog{
			Kita:         int64(b.Kita),
			CreatorId:    int64(b.CreatorId),
			Title:        b.Title,
			Description:  b.Description,
			SystemStatus: convertSystemStatus(b.SystemStatus),
			ImageLinks:   b.ImageLinks,
			CreationDate: timestamppb.New(b.CreationDate),
			//BlogStatuses: make([]*blog.BlogStatus, len(b.BlogStatuses)),
			//BlogRatings:  make([]*blog.BlogRating, len(b.BlogRatings)),
		}

		/*for i, status := range b.BlogStatuses {
			protoBlogStatus := &blog.BlogStatus{
				Id:     int32(status.Id),
				BlogId: int32(status.BlogId),
				Name:   status.Name,
			}
			protoResult.BlogStatuses[i] = protoBlogStatus
		}*/

		protoMessages = append(protoMessages, protoResult)
	}

	return &blog.MultiBlogResponse{Blogs: protoMessages}, nil
}

func MapGrpcBlogToModelBlog(grpcBlog *blog.Blog) *model.Blog {
	modelBlog := &model.Blog{
		Kita:         int(grpcBlog.Kita),
		CreatorId:    int(grpcBlog.CreatorId),
		Title:        grpcBlog.Title,
		Description:  grpcBlog.Description,
		SystemStatus: mapGrpcSystemStatus(grpcBlog.SystemStatus),
		ImageLinks:   grpcBlog.ImageLinks,
		CreationDate: time.Unix(grpcBlog.CreationDate.Seconds, int64(grpcBlog.CreationDate.Nanos)),
	}

	/*for _, grpcStatus := range grpcBlog.BlogStatuses {
		modelStatus := model.BlogStatus{
			Id:     int(grpcStatus.Id),
			BlogId: int(grpcStatus.BlogId),
			Name:   grpcStatus.Name,
		}
		modelBlog.BlogStatuses = append(modelBlog.BlogStatuses, modelStatus)
	}

	for _, grpcRating := range grpcBlog.BlogRatings {
		modelRating := model.BlogRating{
			Id:           int(grpcRating.Id),
			BlogId:       int(grpcRating.BlogId),
			UserId:       int(grpcRating.UserId),
			Rating:       grpcRating.Rating,
			CreationTime: time.Unix(grpcRating.CreationTime.Seconds, int64(grpcRating.CreationTime.Nanos)),
		}
		modelBlog.BlogRatings = append(modelBlog.BlogRatings, modelRating)
	}*/

	return modelBlog
}

func mapGrpcSystemStatus(status blog.Blog_SystemStatus) model.SystemStatus {
	switch status {
	case blog.Blog_DRAFT:
		return model.DRAFT
	case blog.Blog_PUBLISHED:
		return model.PUBLISHED
	case blog.Blog_CLOSED:
		return model.CLOSED
	default:
		return model.DRAFT // Default value if unknown
	}
}

func (handler *BlogHandler) Create(ctx context.Context, request *blog.Blog) (*blog.EmptyResponse, error) {

	newBlog := MapGrpcBlogToModelBlog(request)
	err := handler.BlogService.Create(newBlog)

	if err != nil {
		return nil, err
	}

	return &blog.EmptyResponse{}, nil
}

func (handler *BlogHandler) Update(ctx context.Context, request *blog.Blog) (*blog.BlogResponse, error) {

	updated := MapGrpcBlogToModelBlog(request)
	result, err := handler.BlogService.Update(updated)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	protoResult := &blog.Blog{
		Kita:         int64(result.Kita),
		CreatorId:    int64(result.CreatorId),
		Title:        result.Title,
		Description:  result.Description,
		SystemStatus: convertSystemStatus(result.SystemStatus),
		ImageLinks:   result.ImageLinks,
		CreationDate: timestamppb.New(result.CreationDate),
		//BlogStatuses: make([]*blog.BlogStatus, len(result.BlogStatuses)),
		//BlogRatings:  make([]*blog.BlogRating, len(result.BlogRatings)),
	}

	/*for i, status := range result.BlogStatuses {
		protoBlogStatus := &blog.BlogStatus{
			Id:     int32(status.Id),
			BlogId: int32(status.BlogId),
			Name:   status.Name,
		}
		protoResult.BlogStatuses[i] = protoBlogStatus
	}*/

	/*for j, rating := range result.BlogRatings {
		protoBlogRating := &blog.BlogRating{
			Id:           int32(rating.Id),
			BlogId:       int32(rating.BlogId),
			UserId:       int32(rating.UserId),
			Rating:       rating.Rating,
			CreationTime: timestamppb.New(rating.CreationTime),
		}
		protoResult.BlogRatings[j] = protoBlogRating
	}*/

	return &blog.BlogResponse{Blog: protoResult}, nil
}

func (handler *BlogHandler) Delete(ctx context.Context, request *blog.Id) (*blog.EmptyResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	err := handler.BlogService.Delete(id)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResponse{}, nil

}

/*func (handler *BlogHandler) AddRating(writer http.ResponseWriter, req *http.Request) {
	rating := model.BlogRating{}
	err := json.NewDecoder(req.Body).Decode(&rating)
	if err != nil {
		fmt.Println("lets go")
		//log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.AddRating(&rating)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}*/
