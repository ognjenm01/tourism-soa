package model

import "time"

type SystemStatus int

type Blog struct {
	Id           int
	CreatorId    int
	Title        string
	Description  string
	SystemStatus SystemStatus
	ImageLinks   []string
	CreationDate time.Time
	BlogStatuses []BlogStatus
	BlogRatings  []BlogRating
}
