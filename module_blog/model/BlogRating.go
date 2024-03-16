package model

import "time"

type Rating int

type BlogRating struct {
	BlogId       int
	UserId       int
	Rating       Rating
	CreationTime time.Time
}
