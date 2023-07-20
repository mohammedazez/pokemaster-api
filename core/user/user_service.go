package user

import (
	"context"
	domain "pokemaster-api/core/domain/user"
	port "pokemaster-api/core/port/user"
)

type Service struct {
	repo port.Repository
}

func NewService(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Insert(form *domain.User) error {
	ctx := context.Background()

	err := s.repo.RegisterUser(ctx, form)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) List() ([]domain.User, error) {
	users, err := s.repo.GetAllListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
