package pokemon

import (
	"context"
	domain "pokemaster-api/core/domain/pokemon"
)

type (
	Repository interface {
		InsertPokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error)
		UpdatePokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error)
	}
)
