package pokemon

import (
	"context"
	"errors"
	"log"
	"math/rand"
	domain "pokemaster-api/core/domain/pokemon"
	port "pokemaster-api/core/port/pokemon"
	"strconv"
	"time"
)

var (
	RenameCounter = -1
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

	min := 1
	max := 100
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min+1) + min

	if !isPrime(randomNumber) {
		return domain.Pokemon{}, errors.New("number is not prime")
	}

	result, err := s.repo.InsertPokemon(ctx, form)
	if err != nil {
		return result, err
	}

	return result, err
}

func (s *Service) CatchPokemon() (domain.CatchPokemon, error) {
	var resp domain.CatchPokemon

	rand.Seed(time.Now().UnixNano())
	probability := rand.Float32()
	success := probability <= 0.5

	if !success {
		resp = domain.CatchPokemon{
			Probability: probability,
			Success:     success,
			Information: "failed to catch pokemon",
		}
	} else {
		resp = domain.CatchPokemon{
			Probability: probability,
			Success:     success,
			Information: "success to catch pokemon",
		}
	}

	return resp, nil
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func (s *Service) Update(form *domain.Pokemon) (domain.Pokemon, error) {

	ctx := context.Background()
	RenameCounter++
	fibNum := fibonacci(RenameCounter)

	renamedName := form.PokemonName
	if fibNum >= 0 {
		renamedName = form.PokemonName + "-" + strconv.Itoa(fibNum)
	}

	form.PokemonName = renamedName
	update, err := s.repo.UpdatePokemon(ctx, form)
	if err != nil {
		log.Printf("Failed to update pokemon: %v", err)
		return update, err
	}

	return domain.Pokemon{}, nil
}

func (s *Service) List(pokemonName string, userID string) ([]domain.Pokemon, error) {
	pokemon, err := s.repo.GetAllListPokemon(pokemonName, userID)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
