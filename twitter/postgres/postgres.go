package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/syedwshah/twitter/config"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, conf *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatalf("can't parse postgres config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	db := &DB{Pool: pool}

	db.Ping((ctx))

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("cannot ping postgres: %v", err)
	}

	log.Println("connected to postgres")
}

func (db *DB) Close() {
	db.Pool.Close()
}
