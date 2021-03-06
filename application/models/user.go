package models

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type ListUser struct {
	Username string `json:"username"`
	Image    string `json:"image"`
}
