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
	Kita         int          `bson:"kita" json:"kita" gorm:"primaryKey,unique"`
	CreatorId    int          `bson:"creatorId" json:"creatorId"`
	Title        string       `bson:"title" json:"title"`
	Description  string       `bson:"description" json:"description"`
	SystemStatus SystemStatus `bson:"systemStatus" json:"systemStatus"`
	ImageLinks   string       `bson:"imageLinks" json:"imageLinks"`
	CreationDate time.Time    `bson:"creationDate" json:"creationDate"`
	//BlogStatuses []BlogStatus `bson:"blogStatuses" json:"blogStatuses" gorm:"foreignKey:BlogId;references:Kita"`
	//BlogRatings  []BlogRating `bson:"blogRatings" json:"blogRatings" gorm:"foreignKey:BlogId;references:Kita"`
}

/*func (b *Blog) AddRating(blogRating *BlogRating) {
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
}*/
