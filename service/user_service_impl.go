package service

import (
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Register(req model.CreateUserRequest) (response *model.CreateUserResponse, err error) {
	result, err := service.UserRepository.Register(req)
	if err == nil {

		response = &model.CreateUserResponse{
			AccountNumber: result.NationalIdentityNumber + result.PhoneNumber,
		}
		return response, nil
	} else {
		return nil, err
	}
}
