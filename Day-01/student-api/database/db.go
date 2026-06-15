package database

import (
	"context"


	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(connString string) error {
	var err error

	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}

	return DB.Ping(context.Background())
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}