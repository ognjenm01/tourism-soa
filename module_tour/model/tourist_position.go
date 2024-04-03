package model

import "time"

type TouristPosition struct {
	ID           int          `json:"id" gorm:"primaryKey"`
	TourProgress TourProgress `json:"-" gorm:"foreignKey:TouristPositionId;references:ID"`
	UserId       int          `json:"userid"`
	Latitude     float64      `json:"latitude"`
	Longitude    float64      `json:"longitude"`
	UpdatedAt    time.Time    `json:"updatedat"`
}
