package model

import "time"

type MutationRequest struct {
	AccountNumber string `json:"no_rekening"`
}

type MutationResponses struct {
	MutationDate    time.Time `json:"date"`
	TransactionCode string    `json:"transaction_code"`
	Amount          int64     `json:"amount"`
}
