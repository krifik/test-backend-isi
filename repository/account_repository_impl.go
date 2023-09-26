package repository

import (
	"github/krifik/test-isi/config"
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"

	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	DB *gorm.DB
}

func NewAccountRepositoryImpl(db *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{
		DB: db,
	}
}
func (repository *AccountRepositoryImpl) SavingOrWithdraw(req model.AccountRequest) (*entity.Account, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	accountEntity := &entity.Account{
		AccountNumber: req.AccountNumber,
		Balance:       req.Balance,
	}
	if err := repository.DB.WithContext(ctx).Where("account_number = ?", accountEntity.AccountNumber).Updates(&accountEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	var account entity.Account
	if err := repository.DB.Where("account_number = ?", accountEntity.AccountNumber).First(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := repository.DB.WithContext(ctx).Create(&entity.Mutation{
		AccountID:       account.ID,
		TransactionCode: req.TransactionCode,
		Amount:          req.Amount,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return accountEntity, nil

}

func (repository *AccountRepositoryImpl) FindAccount(req model.AccountRequest) *entity.Account {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var account *entity.Account
	rows := repository.DB.WithContext(ctx).Find(&account, "account_number = ?", req.AccountNumber).RowsAffected
	if rows > 0 {
		return account
	} else {
		return nil
	}
}
