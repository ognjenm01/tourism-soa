package model

import "time"

type TourReview struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	Rating     int         `json:"rating"`
	Comment    string      `json:"comment"`
	VisitDate  time.Time   `json:"visitDate"`
	RatingDate time.Time   `json:"ratingDate"`
	ImageLinks []ImageLink `json:"imageLinks" gorm:"foreignKey:TourReviewID;references:ID"`
	UserId     int         `json:"userId"`
	UserInfo   UserInfo    `json:"userInfo" gorm:"-"`
	TourId     int         `json:"tourId"`
}

type ImageLink struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	TourReviewID int    `json:"tourReviewID"`
	Link         string `json:"link"`
}

type UserInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
