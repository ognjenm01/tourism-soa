package model

import (
	"time"
)

type Tour struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserId           uint      `json:"userid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	Difficulty       int       `json:"difficulty"`
	TransportType    int       `json:"transportType"`
	Status           int       `json:"status"`
	Tags             []TourTag `json:"tags" gorm:"foreignKey:TourID;references:ID"`
	Duration         int       `json:"duration"`
	Distance         float64   `json:"distance"`
	StatusUpdateTime time.Time `json:"statusUpdateTime"`
}

type TourTag struct {
	TourID uint   `json:"TourID"`
	Tag    string `json:"tag"`
}
