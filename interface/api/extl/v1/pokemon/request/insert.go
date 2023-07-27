package request

type (
	RequestInsert struct {
		PokemonName    string `json:"pokemon_name" validate:"required" required:"the name of pokemon is required"`
		PokemonPicture string `json:"pokemon_picture"  validate:"required" required:"the picture of pokemon is required"`
		UserID         string `json:"user_id"  validate:"required" required:"user idis required"`
	}
)
