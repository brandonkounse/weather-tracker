package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ErrMissingEnvFile = "Error Loading .env file"
	ErrNoAPIKey       = "API_KEY is empty"
)

func GetAPIKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ErrMissingEnvFile)
		return "", errors.New(ErrMissingEnvFile)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal(ErrNoAPIKey)
		return "", errors.New(ErrNoAPIKey)
	}

	return apiKey, nil
}
