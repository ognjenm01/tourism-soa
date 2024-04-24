package model

import "time"

type BlogComment struct {
	Id           int       `bson:"_id" json:"id" gorm:"primaryKey"`
	BlogId       int       `bson:"blogId" json:"blogId"`
	UserId       int       `bson:"userId" json:"userId"`
	Comment      string    `bson:"comment" json:"comment"`
	PostTime     time.Time `bson:"postTime" json:"postTime"`
	LastEditTime time.Time `bson:"lastEditTime" json:"lastEditTime"`
	IsDeleted    bool      `bson:"isDeleted" json:"isDeleted"`
}
