package model

type TourEquipment struct {
	ID          int `json:"id" gorm:"primaryKey"`
	TourID      int `json:"tourid"`
	EquipmentID int `json:"equipmentid"`
}
