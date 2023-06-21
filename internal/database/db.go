package database

import (
	"context"
	"fmt"
	"profbuh/internal/config"
	"profbuh/internal/logging"

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

	err = pool.Ping(context.Background())
	if err != nil {
		logging.Log.Fatalf("Can't access db: %v", err)
		return nil, err
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
		password VARCHAR(255) NOT NULL,
		fio VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		logging.Log.Debug("Can't create users table")
		return err
	}

	// TODO: implement main_article_id
	_, err = db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS records (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		video_link VARCHAR(255) NOT NULL,
		status VARCHAR(255) NOT NULL DEFAULT 'В обработке',
		visibility BOOLEAN NOT NULL DEFAULT FALSE,
		author_id INTEGER NOT NULL
	)`)
	if err != nil {
		logging.Log.Debug("Can't create records table")
		return err
	}

	_, err = db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS articles (
		id SERIAL PRIMARY KEY,
		body TEXT NOT NULL,
		record_id INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		logging.Log.Debug("Can't create articles table")
		return err
	}

	// add relation in records table with article id if relation didnt exist
	_, err = db.Pool.Query(context.Background(), `
	ALTER TABLE records
	ADD COLUMN IF NOT EXISTS article_id INTEGER,
	ADD FOREIGN KEY (article_id) REFERENCES articles (id)
	`)
	if err != nil {
		logging.Log.Debug("Can't add relation in records table with article id")
		return err
	}

	// _, err = db.Pool.Query(context.Background(), `
	// ALTER TABLE articles
	// ADD COLUMN IF NOT EXISTS record_id INTEGER,

	_, err = db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS medias (
		id SERIAL PRIMARY KEY,
		link VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		logging.Log.Debug("Can't create medias table")
		return err
	}

	_, err = db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS records_medias (
		record_id INTEGER NOT NULL,
		media_id INTEGER NOT NULL,
		FOREIGN KEY (record_id) REFERENCES records (id),
		FOREIGN KEY (media_id) REFERENCES medias (id)
	)`)
	if err != nil {
		logging.Log.Debug("Can't create records_media table")
		return err
	}

	_, err = db.Pool.Query(context.Background(), `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		comment TEXT NOT NULL,
		record_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (record_id) REFERENCES records (id),
		FOREIGN KEY (author_id) REFERENCES users (id)
	)`)
	if err != nil {
		logging.Log.Debug("Can't create comments table")
		return err
	}

	return nil
}
