package models

type Credit struct {
	UserId    uint    `json:"user_id"`
	AccountNo uint    `json:"account_no"`
	UserName  string  `json:"user_name"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"`
}