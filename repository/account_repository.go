package repository

import (
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"
)

type AccountRepository interface {
	SavingOrWithdraw(req model.AccountRequest) (*entity.Account, error)
	FindAccount(req model.AccountRequest) *entity.Account
}
