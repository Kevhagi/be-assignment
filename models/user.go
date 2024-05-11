package model

import "time"

type User struct {
	Id                string
	Email             string
	Password          string
	SupertokensUserId string
	CreatedAt         time.Time
}
