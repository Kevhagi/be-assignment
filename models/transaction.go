package model

import "time"

type Transaction struct {
	Id                 string
	Type               string
	SourceId           string
	DestinationId      string
	Amount             float32
	Status             string
	IssuedTimestamp    time.Time
	CancelledTimestamp time.Time
	RefusedTimestamp   time.Time
	CompleteTimestamp  time.Time
	Source             Account
	Destination        Account
}
