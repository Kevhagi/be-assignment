package model

import "time"

type Model struct {
	Id           string
	AccountId    string
	Period       string
	DueDate      time.Time
	Amount       float32
	PaidAmount   float32
	UnpaidAmount float32
	CreatedAt    time.Time
	Account      Account
}
