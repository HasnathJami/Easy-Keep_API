package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ReportId    primitive.ObjectID `bson:"report_id"`
	UserId      string             `json:"user_id"`
	UserName    string             `json:"user_name"`
	AccountNo   string             `json:"account_no"`
	TotalAmount string             `json:"total_amount"`
	Date        string             `json:"date"`
}