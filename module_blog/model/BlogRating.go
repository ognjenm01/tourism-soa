package model

import "time"

type Rating int

type BlogRating struct {
	BlogId       int       `json:"blogId"`
	UserId       int       `json:"userId"`
	Rating       Rating    `json:"rating"`
	CreationTime time.Time `json:"creationTime"`
}
