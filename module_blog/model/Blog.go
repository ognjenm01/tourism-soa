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

func (b *Blog) AddRating(blogRating *BlogRating) {
	if b.BlogRatings == nil {
		b.BlogRatings = make([]BlogRating, 0)
	}

	var foundIndex = -1
	for i, rating := range b.BlogRatings {
		if rating.UserId == blogRating.UserId && rating.BlogId == blogRating.BlogId {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		b.BlogRatings[foundIndex] = *blogRating
	} else {
		b.BlogRatings = append(b.BlogRatings, *blogRating)
	}
}
