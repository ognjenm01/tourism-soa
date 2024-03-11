package model

import (
	"time"
)

type TourDifficulty int
type TransportType int
type TourStatus int

type Tour struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	UserId           int            `json:"userid"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Price            float64        `json:"price"`
	Difficulty       TourDifficulty `json:"difficulty"`
	TransportType    TransportType  `json:"transportType"`
	Status           TourStatus     `json:"status"`
	Tags             []TourTag      `json:"tags" gorm:"foreignKey:TourID;references:ID"`
	Keypoints        []Keypoint     `json:"keypoints" gorm:"foreignKey:TourID;references:ID"`
	Duration         int            `json:"duration"`
	Distance         float64        `json:"distance"`
	StatusUpdateTime time.Time      `json:"statusUpdateTime"`
}

type TourTag struct {
	TourID int    `json:"TourID"`
	Tag    string `json:"tag"`
}

func (t *Tour) ContainsStatus(statuses []TourStatus) bool {
	for _, status := range statuses {
		if t.Status == status {
			return true
		}
	}
	return false
}
