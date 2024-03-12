package model

type Keypoint struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	TourID      int     `json:"tourID"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
	Position    uint    `json:"position"`
	Image       string  `json:"image"`
	Secret      string  `json:"secret"`
}
