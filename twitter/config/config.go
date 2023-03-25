package config

import (
	"os"

	"github.com/joho/godotenv"
)

// database is going to be a private struct
type database struct {
	URL string
}

// Config will use a public "Database" of a private "database" struct
type Config struct {
	Database database
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
}
