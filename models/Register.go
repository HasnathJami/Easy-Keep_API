package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Register struct {
	Id       primitive.ObjectID `bson:"id"`
	Name     string             `json:"name"`
	Age      uint               `json:"age"`
	Email    string             `json:"email"`
	Phone    uint               `json:"phone"`
	Username uint               `json:"username"`
	Password string             `json:"password"`
}