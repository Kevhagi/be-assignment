package model

import "time"

type Account struct {
	Id        string
	UserId    string
	Balance   float32
	CreatedAt time.Time
	User      User
}
