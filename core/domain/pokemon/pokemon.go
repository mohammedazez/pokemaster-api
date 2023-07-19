package pokemon

type (
	Pokemon struct {
		ID             string
		PokemonName    string
		PokemonPicture string
		Number         int
		UserID         int
		CreatedAt      string
		UpdatedAt      string
	}

	CatchPokemon struct {
		Probability float32
		Success     bool
		Information string
	}
)
