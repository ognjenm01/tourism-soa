package model

import "time"

type BlogRating struct {
	Id           int       `json:"id"`
	BlogId       int       `json:"blogId"`
	UserId       int       `json:"userId"`
	Rating       string    `json:"rating"`
	CreationTime time.Time `json:"creationTime"`
}
