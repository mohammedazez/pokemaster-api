package response

import (
	"pokemaster-api/interface/api/common"
)

type (
	Response struct {
		ID             string `json:"id"`
		PokemonName    string `json:"pokemon_name"`
		PokemonPicture string `json:"pokemon_picture"`
		Number         int    `json:"number"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	}

	PrimeNumberCheck struct {
		Result      string `json:"result"`
		Released    bool   `json:"released"`
		PrimeNumber int    `json:"prime_number"`
	}

	CatchPokemon struct {
		Success     bool    `json:"success"`
		Probability float32 `json:"probability"`
		Information string  `json:"information"`
	}
)

func NewResponseDetail(message string, res *PrimeNumberCheck, code int, status bool) *common.DefaultResponse {
	data := new(PrimeNumberCheck)
	data.Result = res.Result
	data.Released = res.Released
	data.PrimeNumber = res.PrimeNumber

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, data, code, status)
	return responseData
}

func NewResponseCatchPokemon(message string, res *CatchPokemon, code int, status bool) *common.DefaultResponse {
	data := new(CatchPokemon)
	data.Probability = res.Probability
	data.Success = res.Success
	data.Information = res.Information

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, data, code, status)
	return responseData
}
