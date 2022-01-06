package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rhuancaetano/enceladus/pkg/envalid"
)

type Environment struct {
	PORT              string `env:"default:5000"`
	POSTGRES_USER     string `env:"required"`
	POSTGRES_PASSWORD string `env:"required"`
	POSTGRES_DB       string `env:"required"`
	POSTGRES_HOST     string `env:"required"`
	JWT_SECRET        string `env:"required"`
}

var environment = &Environment{}

func GetEnvironment() *Environment {
	if env, ok := os.LookupEnv("ENV"); ok {
		if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
			log.Fatalf("Error loading .env.%s file", env)
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	envalid.GetEnvironmentVariables(environment, os.LookupEnv)

	return environment
}
