package models

type Report struct {
	ReportId    string `json:"report_id"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	AccountNo   string `json:"account_no"`
	TotalAmount string `json:"total_amount"`
	Date        string `json:"date"`
}