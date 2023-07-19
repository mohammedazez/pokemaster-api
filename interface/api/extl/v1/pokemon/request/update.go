package request

type (
	RequestUpdate struct {
		ID          string `json:"-"`
		PokemonName string `json:"pokemon_name" validate:"required" required:"the name of pokemon is required"`
	}
)
