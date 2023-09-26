package repository

import (
	"fmt"
	"github/krifik/test-isi/config"
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"

	"gorm.io/gorm"
)

type MutationRepositoryImpl struct {
	DB *gorm.DB
}

func NewMutationRepositoryImpl(db *gorm.DB) MutationRepository {
	return &MutationRepositoryImpl{
		DB: db,
	}
}

func (repository *MutationRepositoryImpl) FindMutations(req model.MutationRequest) *[]entity.Mutation {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	accountRepository := &AccountRepositoryImpl{
		DB: repository.DB,
	}
	fmt.Println("test", req)
	accountEntity := accountRepository.FindAccount(model.AccountRequest{
		AccountNumber: req.AccountNumber,
	})
	var mutation *[]entity.Mutation
	if tx := repository.DB.WithContext(ctx).Find(&mutation, "account_id = ?", accountEntity.ID).Error; tx != nil {
		return nil
	} else {
		return mutation
	}
}
