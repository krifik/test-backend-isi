package config

import "github/krifik/test-isi/entity"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entity.User{}},
		{Entity: entity.Account{}},
		{Entity: entity.Mutation{}},
	}
}
