package model

type BlogStatus struct {
	Id     int    `json:"id"`
	BlogId int    `json:"blogId"`
	Name   string `json:"name"`
}
