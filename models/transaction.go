package models

type Transaction struct {
	Id       int    `json:"id" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
	Balance  string `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
	Category string `json:"category"`
}
