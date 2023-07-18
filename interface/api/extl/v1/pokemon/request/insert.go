package request

type (
	RequestInsert struct {
		PokemonName    string `json:"pokemon_name" validate:"required" required:"the name of pokemon is required"`
		PokemonPicture string `json:"pokemon_picture"  validate:"required" required:"the picture of pokemon is required"`
		Number         int    `json:"number"  validate:"required" required:"the number is required"`
	}
)
