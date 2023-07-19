package response

type (
	Response struct {
		ID             string `json:"id"`
		PokemonName    string `json:"pokemon_name"`
		PokemonPicture string `json:"pokemon_picture"`
		Number         int    `json:"number"`
		UserID         int    `json:"user_id"`
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

	UpdatePokemon struct {
		ID          string `json:"id"`
		PokemonName string `json:"pokemon_name"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
)
