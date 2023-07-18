package pokemon

import (
	"context"
	domain "pokemaster-api/core/domain/pokemon"
	"pokemaster-api/infrastructure/repository/mysql/transactor"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Pokemon struct {
		ID             string `gorm:"primaryKey;column:id"`
		PokemonName    string `gorm:"column:pokemon_name"`
		PokemonPicture string `gorm:"column:pokemon_picture"`
		PrimeNumber    int    `gorm:"column:prime_number"`
		CreatedAt      string `gorm:"column:created_at"`
		UpdatedAt      string `gorm:"column:updated_at"`
	}
)

func (Pokemon) TableName() string {
	return "pokemon"
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) getDB(ctx context.Context) *gorm.DB {
	dbWithTx := transactor.ExtractTx(ctx)
	if dbWithTx != nil {
		return dbWithTx
	}
	return repo.db
}

func (repo *Repository) InsertPokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error) {
	pokemons := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(&pokemons).Create(&pokemons).Error; err != nil {
		return domain.Pokemon{}, err
	}

	outData := new(domain.Pokemon)
	outData.ID = pokemons.ID
	outData.Number = pokemons.PrimeNumber

	return *outData, nil
}

func mappingInput(pokemon *domain.Pokemon) Pokemon {
	var result Pokemon

	id := uuid.New()

	timeNow := time.Now()
	result.ID = id.String()
	result.PokemonName = pokemon.PokemonName
	result.PokemonPicture = pokemon.PokemonPicture
	result.PrimeNumber = pokemon.Number
	result.CreatedAt = timeNow.String()

	return result
}
