package model

type TourTag struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	TourID int    `json:"TourID"`
	Tag    string `json:"tag"`
}
