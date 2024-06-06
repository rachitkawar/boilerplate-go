package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rachitkawar/boilerplate-go/src/common"
	"os"
)

type DB struct {
	db *pgxpool.Pool
}

func InitDB(ctx context.Context) *DB {
	var err error
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		common.Log.Fatal("Unable to connect to database: %v\n", err)
	}

	return &DB{db: db}
}
