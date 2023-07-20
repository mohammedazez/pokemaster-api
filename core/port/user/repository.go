package user

import (
	"context"
	domain "pokemaster-api/core/domain/user"
)

type (
	Repository interface {
		RegisterUser(ctx context.Context, inData *domain.User) error
		GetAllListUsers() ([]domain.User, error)
	}
)
