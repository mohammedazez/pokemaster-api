package request

type (
	RequestList struct {
		PokemonName string `query:"pokemon_name"`
	}
)
