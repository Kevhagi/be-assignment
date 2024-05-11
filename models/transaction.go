package model

import "time"

type Transaction struct {
	Id            string
	Type          string
	SourceId      string
	DestinationId string
	Amount        float32
	Currency      string
	Status        string
	Timestamp     time.Time
	Source        Account
	Destination   Account
}
