package repository

import (
	"github/krifik/test-isi/entity"
	"github/krifik/test-isi/model"
)

type MutationRepository interface {
	FindMutations(req model.MutationRequest) *[]entity.Mutation
}
