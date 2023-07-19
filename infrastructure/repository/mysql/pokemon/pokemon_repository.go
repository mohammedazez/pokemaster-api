package pokemon

import (
	"context"
	"errors"
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
		UserID         int    `gorm:"column:user_id"`
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

func (repo *Repository) UpdatePokemon(ctx context.Context, inData *domain.Pokemon) (domain.Pokemon, error) {
	pokemons := mappingInput(inData)

	db := repo.getDB(ctx)
	timeNow := time.Now()
	pokemons.UpdatedAt = timeNow.String()

	updates := map[string]interface{}{
		"pokemon_name": pokemons.PokemonName,
		"updated_at":   pokemons.UpdatedAt,
	}

	err := db.Model(&Pokemon{}).Where("id = ?", inData.ID).Updates(updates).Error
	if err != nil {
		return domain.Pokemon{}, err
	}

	outData := new(domain.Pokemon)
	outData.ID = pokemons.ID
	outData.PokemonName = pokemons.PokemonName
	outData.CreatedAt = pokemons.CreatedAt
	outData.UpdatedAt = pokemons.UpdatedAt

	return *outData, nil
}

func (repo *Repository) GetPokemon(ctx context.Context, ID string) (domain.Pokemon, error) {
	var pokemon domain.Pokemon
	db := repo.getDB(ctx)

	err := db.Table("pokemon").First(&pokemon, "id = ?", ID).Error
	if err != nil {
		return domain.Pokemon{}, err
	}
	return pokemon, nil
}

func (repo *Repository) GetAllListPokemon(pokemonName string) ([]domain.Pokemon, error) {
	pokemons := make([]Pokemon, 0)
	query := repo.db.Table("pokemon")

	if pokemonName != "" {
		query = query.Where("pokemon_name LIKE ?", "%"+pokemonName+"%")
	}

	result := query.Find(&pokemons)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}

	outData := make([]domain.Pokemon, 0)
	for _, value := range pokemons {
		var pokemon domain.Pokemon
		pokemon.ID = value.ID
		pokemon.PokemonName = value.PokemonName
		pokemon.PokemonPicture = value.PokemonPicture
		pokemon.UserID = value.UserID
		pokemon.Number = value.PrimeNumber
		pokemon.CreatedAt = value.CreatedAt
		pokemon.UpdatedAt = value.UpdatedAt
		outData = append(outData, pokemon)
	}

	return outData, nil
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
