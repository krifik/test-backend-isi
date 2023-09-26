package service

import (
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/repository"
)

type MutationServiceImpl struct {
	MutationRepository repository.MutationRepository
}

// FindMutations implements MutationService
func (service *MutationServiceImpl) FindMutations(req model.MutationRequest) *[]entity.Mutation {
	result := service.MutationRepository.FindMutations(req)
	if result != nil {
		return result
	} else {
		return nil
	}
}

func NewMutationServiceImpl(mutationRepository repository.MutationRepository) MutationService {
	return &MutationServiceImpl{
		MutationRepository: mutationRepository,
	}
}
