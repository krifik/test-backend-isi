package repository

import (
	"errors"
	"github/krifik/test-isi/config"
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Register(req model.CreateUserRequest) (user *entity.User, err error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	user = &entity.User{
		Name:                   req.Name,
		NationalIdentityNumber: req.NationalIdentityNumber,
		PhoneNumber:            req.PhoneNumber,
	}

	isExist := repository.FindUser(req)

	if isExist {
		return nil, errors.New("user already exist")
	}

	if err = repository.DB.WithContext(ctx).Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = repository.DB.WithContext(ctx).Create(&entity.Account{
		UserID:        user.ID,
		AccountNumber: user.NationalIdentityNumber + user.PhoneNumber,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) FindUser(req model.CreateUserRequest) bool {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var user *entity.User
	rows := repository.DB.WithContext(ctx).Find(&user, "national_identity_number = ? AND phone_number = ?", req.NationalIdentityNumber, req.PhoneNumber).RowsAffected

	if rows > 0 {
		return true
	} else {
		return false
	}
}
