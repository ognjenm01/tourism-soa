package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"module_blog.xws.com/model"
	"module_blog.xws.com/repo"
)

type BlogService struct {
	BlogRepo          *repo.BlogRepository
	BlogRatingService *BlogRatingService
	BlogStatusService *BlogStatusService
}

func (service *BlogService) GetByPeopleUFollow(id string) (*[]model.Blog, error) {
	followerIds, err := service.GetFollowerIDs(id)

	if err != nil {
		return nil, err
	}
	blogs, err1 := service.BlogRepo.GetByPeopleUFollow(*followerIds)

	if err1 != nil {
		return nil, err1
	}

	return &blogs, nil
}

func (service *BlogService) GetFollowerIDs(id string) (*[]int, error) {
	var people []model.Person
	var ids []int

	resp, err := http.Get(fmt.Sprintf("http://localhost:6666/api/person/followers/%s", id))

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("nema")
		return nil, err
	}

	if err := json.Unmarshal(body, &people); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	log.Fatalln(len(people))

	for _, p := range people {
		ids = append(ids, p.UserId)
	}

	return &ids, nil

}

func (service *BlogService) GetById(id string) (*model.Blog, error) {
	blog, err := service.BlogRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &blog, err

}

/*func (service *BlogService) GetAll() (*[]model.Blog, error) {
	blogs, error := service.BlogRepo.GetAll()

	if error != nil {
		log.Fatalf("Database is empty\n")
		return &[]model.Blog{}, error
	}

	return &blogs, nil
}*/

func (service *BlogService) Create(blog *model.Blog) error {
	result := service.BlogRepo.Create(blog)
	if result != nil {
		return result
	}

	return nil
}

func (service *BlogService) Update(blog *model.Blog) error {
	result := service.BlogRepo.Update(blog)
	if result != nil {
		return result
	}

	return nil
}

func (service *BlogService) Delete(id string) error {
	result := service.BlogRepo.Delete(id)
	if result != nil {
		return result
	}

	return nil
}

func (service *BlogService) updateStatuses(blog model.Blog, status *model.BlogStatus) {
	exists := false

	for _, s := range blog.BlogStatuses {
		if s.Name == status.Name {
			exists = true
			return
		}
	}

	if !exists {
		err3 := service.BlogStatusService.Create(status)

		if err3 != nil {
			log.Fatal(err3)
		}
	}
}

func (service *BlogService) determineStatus(rating *model.BlogRating, blog model.Blog) {
	upvotes := 0
	downvotes := 0
	commentNum := 3
	status := &model.BlogStatus{
		BlogId: rating.BlogId,
	}

	for _, r := range blog.BlogRatings {
		if r.Rating == "UPVOTE" {
			upvotes++
		}
		if r.Rating == "DOWNVOTE" {
			downvotes++
		}
	}
	score := (upvotes - downvotes) * (commentNum + 1)
	if score > 35 && commentNum >= 2 {
		(*status).Name = "POPULAR"
	} else if score > 2 && commentNum >= 1 {
		(*status).Name = "ACTIVE"
	} else if score < 0 {
		(*status).Name = "NEW" //CLOSED
	} else {
		(*status).Name = "NEW"
	}

	service.updateStatuses(blog, status)
}

func (service *BlogService) AddRating(rating *model.BlogRating) error {

	blog, err1 := service.BlogRepo.GetById(strconv.Itoa(rating.BlogId))
	if err1 != nil {
		fmt.Println("1")
	}

	if blog.SystemStatus == model.CLOSED {
		return nil
	}

	blog.AddRating(rating)
	service.determineStatus(rating, blog)

	return nil
}
