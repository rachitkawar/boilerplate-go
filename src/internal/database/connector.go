package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rachitkawar/boilerplate-go/src/utils"
)

type DB struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func InitDB(ctx context.Context) *DB {
	var err error
	db, err := pgxpool.New(ctx, utils.GetEnv("DSN"))
	if err != nil {
		utils.Log.Fatal("Unable to connect to database: %v\n", err)
	}

	return &DB{db: db, ctx: ctx}
}
