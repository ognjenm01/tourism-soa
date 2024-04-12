package model

import "time"

type BlogRating struct {
	Id           int       `bson:"_id" json:"id"`
	BlogId       int       `bson:"blogId" json:"blogId"`
	UserId       int       `bson:"userId" json:"userId"`
	Rating       string    `bson:"rating" json:"rating"`
	CreationTime time.Time `bson:"creationTime" json:"creationTime"`
}
