package models

type Register struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Phone    uint   `json:"phone"`
	Username uint   `json:"username"`
	Password string `json:"password"`
}