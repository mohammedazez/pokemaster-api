package pokemon

type (
	Pokemon struct {
		ID             string
		PokemonName    string
		PokemonPicture string
		Number         int
		UserID         string
		RenameCounter  int
		CreatedAt      string
		UpdatedAt      string
	}

	CatchPokemon struct {
		Probability float32
		Success     bool
		Information string
	}
)
