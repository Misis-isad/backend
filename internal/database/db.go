package database

import (
	"context"
	"fmt"
	"profbuh/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Pool *pgxpool.Pool
}

func InitDb(config *config.Config) (*Db, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbName)
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	db := &Db{
		Pool: pool,
	}

	err = db.CreateTables()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Db) CreateTables() error {
	_, err := db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		return err
	}

	return nil
}
