package pokemon

import (
	"context"
	"errors"
	"log"
	"math/rand"
	domain "pokemaster-api/core/domain/pokemon"
	port "pokemaster-api/core/port/pokemon"
	"strconv"
	"strings"
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

func afters(n int) (after int) {
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	after = -1
	for i := 0; i < len(fib); i++ {
		if fib[i] == n {
			if i < len(fib)-1 {
				after = fib[i+1]
			}
			break
		}
	}
	return after
}

func (s *Service) Update(form *domain.Pokemon) (domain.Pokemon, error) {
	ctx := context.Background()
	getPokemon, err := s.repo.GetPokemon(ctx, form.ID)
	if err != nil {
		log.Printf("Failed to get pokemon: %v", err)
		return getPokemon, err
	}
	var n string
	index := strings.LastIndex(getPokemon.PokemonName, "-")
	for i := index + 1; i < len(getPokemon.PokemonName); i++ {
		n += string(getPokemon.PokemonName[i])
	}
	nInt, _ := strconv.Atoi(n)

	if nInt <= 5 {
		RenameCounter++
		fibNum := fibonacci(RenameCounter)

		renamedName := form.PokemonName
		if fibNum >= 0 {
			renamedName = form.PokemonName + "-" + strconv.Itoa(fibNum)
		}
		form.PokemonName = renamedName
	} else {
		after := afters(nInt)
		form.PokemonName = form.PokemonName + "-" + strconv.Itoa(after)
	}
	update, err := s.repo.UpdatePokemon(ctx, form)
	if err != nil {
		log.Printf("Failed to update pokemon: %v", err)
		return update, err
	}

	return update, nil
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
