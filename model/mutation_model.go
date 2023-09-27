package model

import "time"

type MutationRequest struct {
	AccountNumber string `json:"no_rekening"`
}

type MutationResponses struct {
	MutationDate    time.Time `json:"waktu"`
	TransactionCode string    `json:"kode_transaksi"`
	Amount          int64     `json:"nominal"`
}
