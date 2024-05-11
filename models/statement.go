package model

import "time"

type Statement struct {
	Id            string
	Type          string
	SourceId      string
	DestinationId string
	Amount        float32
	CreatedAt     time.Time
	Source        Account
	Destination   Account
}
