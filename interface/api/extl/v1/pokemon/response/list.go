package response

import (
	domain "pokemaster-api/core/domain/pokemon"
	"pokemaster-api/interface/api/common"
)

type (
	ResponseList struct {
		Response
	}
)

func NewResponseList(message string, data []domain.Pokemon, code int) *common.DefaultResponse {
	pokemons := make([]ResponseList, 0)

	for _, val := range data {
		var pokemon ResponseList

		pokemon.ID = val.ID
		pokemon.PokemonName = val.PokemonName
		pokemon.PokemonPicture = val.PokemonPicture
		pokemon.Number = val.Number
		pokemon.UserID = val.UserID
		pokemon.UpdatedAt = val.UpdatedAt
		pokemons = append(pokemons, pokemon)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, pokemons, code, true)
	return responseData
}
