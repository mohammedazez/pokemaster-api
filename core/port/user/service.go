package user

import (
	domain "pokemaster-api/core/domain/user"
)

type (
	Service interface {
		Insert(form *domain.User) error
		List() ([]domain.User, error)
	}
)
