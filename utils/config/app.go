package config

import (
	"pokemaster-api/utils/config/mysql"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()

	mysql.Init()
}
