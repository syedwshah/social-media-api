package main

import (
	"context"
	"fmt"
	"log"

	"github.com/syedwshah/twitter/config"
	"github.com/syedwshah/twitter/postgres"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("WORKING")
}
