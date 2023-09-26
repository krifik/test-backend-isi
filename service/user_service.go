package service

import (
	"github/krifik/test-isi/model"
)

type UserService interface {
	Register(req model.CreateUserRequest) (response *model.CreateUserResponse, err error)
}
