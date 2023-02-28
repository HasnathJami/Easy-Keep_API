package models

type Login struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}