package model

type Register struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	ProfileImage string `json:"profile_image"`
	Biography    string `json:"biography"`
	Quote        string `json:"quote"`
}
