package service

import (
	"github/krifik/test-isi/model"
)

type MutationService interface {
	FindMutations(req model.MutationRequest) []model.MutationResponses
}
