package service

import "github/krifik/test-isi/model"

type AccountService interface {
	SavingOrWithdraw(req model.AccountRequest, isSaving bool) (response *model.AccountResponse, err error)
	FindAccount(req model.AccountRequest) (response *model.AccountResponse)
}
