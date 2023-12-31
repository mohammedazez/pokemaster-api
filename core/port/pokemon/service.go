package pokemon

import (
	domain "pokemaster-api/core/domain/pokemon"
)

type (
	Service interface {
		Insert(form *domain.Pokemon) (domain.Pokemon, error)
		CatchPokemon() (domain.CatchPokemon, error)
		Update(form *domain.Pokemon) (domain.Pokemon, error)
		List(pokemonName string, userID string) ([]domain.Pokemon, error)
	}
)
