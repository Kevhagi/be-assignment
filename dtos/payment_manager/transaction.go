package paymentmanagerdto

import (
	"time"
)

type TransactionSendRequestController struct {
	DestinationAccountID string  `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
	Currency             string  `json:"currency"`
}

type TransactionSendRequestRepository struct {
	SourceAccountID      string  `json:"source_account_id"`
	DestinationAccountID string  `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
	Currency             string  `json:"currency"`
}

type TransactionSendResponseRepository struct {
	TransactionID        string    `json:"transaction_id"`
	Timestamp            time.Time `json:"timestamp"`
	SourceAccountID      string    `json:"source_account_id"`
	DestinationAccountID string    `json:"destination_account_id"`
}
