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

type jwt struct {
	Secret string
	Issuer string //service issuer
}

type Config struct {
	Database database // public "Database" of a private "database" struct
	JWT      jwt
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
		JWT: jwt{
			Secret: os.Getenv("JWT_SECRET"),
			Issuer: os.Getenv("DOMAIN"),
		},
	}
}
