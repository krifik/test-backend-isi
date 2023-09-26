package model

type AccountRequest struct {
	AccountNumber   string `json:"no_rekening"`
	Amount          int64  `json:"nominal"`
	Balance         int64  `json:"saldo"`
	TransactionCode string `json:"-"`
}

type AccountResponse struct {
	Balance int64 `json:"saldo"`
}
