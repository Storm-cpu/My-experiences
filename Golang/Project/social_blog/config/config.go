package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDB       string
	PostgresSSLMode  string
	PostgresPassword string
}

func LoadConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Configuration{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresSSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}

	return config
}
