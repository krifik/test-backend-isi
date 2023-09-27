package service

import (
	"github/krifik/test-isi/model"
	"github/krifik/test-isi/repository"
)

type MutationServiceImpl struct {
	MutationRepository repository.MutationRepository
}

// FindMutations implements MutationService
func (service *MutationServiceImpl) FindMutations(req model.MutationRequest) []model.MutationResponses {
	result := service.MutationRepository.FindMutations(req)
	if result != nil {
		var response []model.MutationResponses
		for _, result := range *result {
			response = append(response, model.MutationResponses{
				MutationDate:    result.CreatedAt,
				TransactionCode: result.TransactionCode,
				Amount:          result.Amount,
			})
		}
		return response
	} else {
		return nil
	}
}

func NewMutationServiceImpl(mutationRepository repository.MutationRepository) MutationService {
	return &MutationServiceImpl{
		MutationRepository: mutationRepository,
	}
}
