package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	return Config{
		Port: port,
	}
}
