package model

import (
	"time"
)

type SystemStatus int

const (
	DRAFT = iota
	PUBLISHED
	CLOSED
)

type Blog struct {
	Id           int          `json:"id" gorm:"primaryKey,unique"`
	CreatorId    int          `json:"creatorId"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	SystemStatus SystemStatus `json:"systemStatus"`
	ImageLinks   string       `json:"imageLinks"`
	CreationDate time.Time    `json:"creationDate"`
	BlogStatuses []BlogStatus `json:"blogStatuses" gorm:"foreignKey:BlogId;references:Id"`
	BlogRatings  []BlogRating `json:"blogRatings" gorm:"foreignKey:BlogId;references:Id"`
}
