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

	Port         int
	ReadTimeout  int
	WriteTimeout int

	JwtAdminSecret    string
	JwtAdminDuration  int
	JwtAdminAlgorithm string
}

func LoadConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %+v", err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresSSLMode := os.Getenv("POSTGRES_SSLMODE")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))

	jwtAdminSecret := os.Getenv("JWT_SECRET")
	jwtAdminDuration, _ := strconv.Atoi(os.Getenv("JWT_DURATION"))
	jwtAdminAlgorithm := os.Getenv("JWT_ALGORITHM")

	config := Configuration{
		PostgresHost:     postgresHost,
		PostgresUser:     postgresUser,
		PostgresDB:       postgresDB,
		PostgresSSLMode:  postgresSSLMode,
		PostgresPassword: postgresPassword,

		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,

		JwtAdminSecret:    jwtAdminSecret,
		JwtAdminDuration:  jwtAdminDuration,
		JwtAdminAlgorithm: jwtAdminAlgorithm,
	}

	return config
}
