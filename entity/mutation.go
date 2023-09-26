package entity

import "gorm.io/gorm"

const (
	WITHDRAW = "D"
	SAVING   = "C"
)

type Mutation struct {
	gorm.Model
	AccountID       int64
	TransactionCode string
	Amount          int64
}
