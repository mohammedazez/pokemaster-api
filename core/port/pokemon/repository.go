package pokemon

import (
	"context"
	domain "pokemaster-api/core/domain/pokemon"
)

type (
	Repository interface {
		InsertPokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error)
		UpdatePokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error)
		GetPokemon(ctx context.Context, ID string) (domain.Pokemon, error)
		GetAllListPokemon(pokemonName string) ([]domain.Pokemon, error)
	}
)
