package model

type TourDifficulty int
type TransportType int
type UserRole int

type User struct {
	ID        int      `json:"id" gorm:"primaryKey,unique"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
	IsActive  bool     `json:"isactive"`
	IsBlocked bool     `json:"isblocked"`
	IsEnabled bool     `json:"isenabled"`
}
