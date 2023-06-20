package service

import (
	"context"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *pgxpool.Pool, c context.Context, userData models.UserCreate) (models.UserDto, error) {
	var err error
	userData.Password, err = HashPassword(userData.Password)
	if err != nil {
		return models.UserDto{}, err
	}

	user, err := crud.CreateUser(db, c, userData)
	if err != nil {
		return models.UserDto{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return password, err
	}

	return string(hash), nil
}
