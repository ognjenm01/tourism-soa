package model

import (
	"encoding/json"
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

func (t *TourDifficulty) UnmarshalJSON(bytes []byte) error {
	var difficulty string
	err := json.Unmarshal(bytes, &difficulty)
	if err != nil {
		return err
	}
	switch difficulty {
	case "EASY":
		*t = 0
	case "MEDIUM":
		*t = 1
	case "HARD":
		*t = 2
	case "EXTREME":
		*t = 3
	}
	return err
}

func (t *TourStatus) UnmarshalJSON(bytes []byte) error {
	var status string
	err := json.Unmarshal(bytes, &status)
	if err != nil {
		return err
	}
	switch status {
	case "DRAFT":
		*t = 0
	case "PUBLISHED":
		*t = 1
	case "ARCHIVED":
		*t = 2
	case "DISABLED":
		*t = 3
	case "CUSTOM":
		*t = 4
	case "CAMPAIGN":
		*t = 5
	}
	return err
}

func (t *TransportType) UnmarshalJSON(bytes []byte) error {
	var transportType string
	err := json.Unmarshal(bytes, &transportType)
	if err != nil {
		return err
	}
	switch transportType {
	case "WALK":
		*t = 0
	case "CAR":
		*t = 1
	case "BIKE":
		*t = 2
	case "BOAT":
		*t = 3
	}
	return err
}
