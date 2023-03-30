//go:build integration
// +build integration

package domain

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/syedwshah/twitter"
	"github.com/syedwshah/twitter/config"
	"github.com/syedwshah/twitter/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf        *config.Config
	db          *postgres.DB
	userRepo    twitter.UserRepo
	authService twitter.AuthService
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	config.LoadEnv(".env.test")

	passwordCost = bcrypt.MinCost

	conf = config.New()

	db = postgres.New(ctx, conf)
	defer db.Close()

	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	userRepo = postgres.NewUserRepo(db)

	authService = NewAuthService(userRepo)

	os.Exit(m.Run())
}
