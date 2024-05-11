package model

import "time"

type Profile struct {
	Id          string
	UserId      string
	FirstName   string
	LastName    string
	PhoneNumber string
	CreatedAt   time.Time
	User        User
}
