package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDB       string
	PostgresSSLMode  string
	PostgresPassword string
	Port             int
	ReadTimeout      int
	WriteTimeout     int
}

func LoadConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %+v", err)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))

	config := Configuration{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresSSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		Port:             port,
		ReadTimeout:      readTimeout,
		WriteTimeout:     writeTimeout,
	}

	return config
}
