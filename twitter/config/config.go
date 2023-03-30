package config

import (
	"os"
	"regexp"

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

func LoadEnv(fileName string) {
	re := regexp.MustCompile(`^(.*` + "twitter" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/` + fileName)
	if err != nil {
		godotenv.Load()
	}
}

func New() *Config {
	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
}
