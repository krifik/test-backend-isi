package service

import (
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/repository"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
}

func NewAccountServiceImpl(accountRepository repository.AccountRepository) AccountService {
	return &AccountServiceImpl{
		AccountRepository: accountRepository,
	}
}

// SavingOrWithdraw implements AccountService
func (service *AccountServiceImpl) SavingOrWithdraw(req model.AccountRequest, isSaving bool) (response *model.AccountResponse, err error) {

	isAccountExist := service.AccountRepository.FindAccount(req)

	if isAccountExist != nil {
		if isSaving {
			req.Balance = isAccountExist.Balance + req.Amount
			req.TransactionCode = entity.SAVING
		} else {
			req.Balance = isAccountExist.Balance - req.Amount
			req.TransactionCode = entity.WITHDRAW
		}
	}
	response = &model.AccountResponse{
		Balance: req.Balance,
	}

	_, err = service.AccountRepository.SavingOrWithdraw(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *AccountServiceImpl) FindAccount(req model.AccountRequest) (response *model.AccountResponse) {
	result := service.AccountRepository.FindAccount(req)
	if result != nil {
		response = &model.AccountResponse{
			Balance: result.Balance,
		}
		return response
	} else {
		return nil
	}
}
