package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT              string `env:"default:5000"`
	POSTGRES_USER     string `env:"required"`
	POSTGRES_PASSWORD string `env:"required"`
	POSTGRES_DB       string `env:"required"`
	JWT_SECRET        string `env:"required"`
}

var environment *Environment

func GetEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if environment == nil {
		environment = &Environment{
			PORT:              os.Getenv("PORT"),
			POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
			JWT_SECRET:        os.Getenv("JWT_SECRET"),
		}
	}

	return environment
}
