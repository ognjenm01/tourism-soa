package model

import "time"

type TourProgressStatus int

type TourProgress struct {
	ID                int                `json:"id" gorm:"primaryKey"`
	TouristPositionId int                `json:"touristpositionid"`
	TourId            int                `json:"tourid"`
	Status            TourProgressStatus `json:"status"`
	StartTime         time.Time          `json:"starttime"`
	LastActivity      time.Time          `json:"lastactivity"`
	CurrentKeypoint   int                `json:"currentkeypoint"`
	IsInProgress      bool               `json:"isinprogress"`
}
