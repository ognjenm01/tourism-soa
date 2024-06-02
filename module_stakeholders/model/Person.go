package model

type Person struct {
	ID           int    `json:"id" gorm:"primaryKey,unique"`
	UserId       int    `json:"userid"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileimage"`
	Biography    string `json:"biography"`
	Quote        string `json:"quote"`
}
