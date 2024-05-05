package model

type Person struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}
