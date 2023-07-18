package pokemon

import (
	"context"
	domain "pokemaster-api/core/domain/pokemon"
	port "pokemaster-api/core/port/pokemon"
)

type Service struct {
	repo port.Repository
}

func NewService(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Insert(form *domain.Pokemon) (domain.Pokemon, error) {
	ctx := context.Background()

	result, err := s.repo.InsertPokemon(ctx, form)
	if err != nil {
		return result, err
	}

	return result, err
}
