package repository

import (
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"
)

type UserRepository interface {
	Register(req model.CreateUserRequest) (user *entity.User, err error)
}
