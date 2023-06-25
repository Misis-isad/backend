package database

import (
	"fmt"
	"profbuh/internal/config"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func InitDb(cfg *config.Config) (*Database, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logging.Log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Record{}, &models.Article{}, &models.Comment{} /*, &models.Media{}*/)
	if err != nil {
		logging.Log.Fatalf("Error migrating database: %v", err)
		return nil, err
	}

	return &Database{Db: db}, nil
}
