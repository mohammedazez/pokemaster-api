package pokemon

import (
	"context"
	"log"
	"math/rand"
	domain "pokemaster-api/core/domain/pokemon"
	port "pokemaster-api/core/port/pokemon"
	"strconv"
	"strings"
	"time"
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

func (s *Service) Update(form *domain.Pokemon) (domain.Pokemon, error) {
	ctx := context.Background()

	getPokemon, err := s.repo.GetPokemon(ctx, form.ID)
	if err != nil {
		log.Printf("Failed to get pokemon: %v", err)
		return getPokemon, err
	}

	fibNum := getNumber(getPokemon.PokemonName)
	form.PokemonName = form.PokemonName + "-" + fibNum

	update, err := s.repo.UpdatePokemon(ctx, form)
	if err != nil {
		log.Printf("Failed to update pokemon: %v", err)
		return update, err
	}

	return update, nil
}

func getNumber(name string) string {
	var n string

	index := strings.LastIndex(name, "-")
	for i := index + 1; i < len(name); i++ {
		n += string(name[i])
	}

	num, _ := strconv.Atoi(n)
	f := fibonacci()
	fibSlice := make([]int, num)
	for i := 0; i < num; i++ {
		fibSlice[i] = f()
	}

	filteredSlice := make([]int, 0)
	for _, value := range fibSlice {
		if value <= num && value >= 0 {
			filteredSlice = append(filteredSlice, value)
		}
	}

	sum := 0
	if len(filteredSlice) >= 2 {
		sum = filteredSlice[len(filteredSlice)-1] + filteredSlice[len(filteredSlice)-2]
	}
	return strconv.Itoa(sum)
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}
