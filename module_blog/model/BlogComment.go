package model

import "time"

type BlogComment struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	BlogId       int       `json:"blogId"`
	UserId       int       `json:"userId"`
	Comment      string    `json:"comment"`
	PostTime     time.Time `json:"postTime"`
	LastEditTime time.Time `json:"lastEditTime"`
	IsDeleted    bool      `json:"isDeleted"`
}
