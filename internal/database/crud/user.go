package crud

import (
	"context"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(db *pgxpool.Pool, ctx context.Context, userData models.UserCreate) (models.UserDto, error) {
	var user models.UserDto

	res := db.QueryRow(ctx, `
		INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email
	`, userData.Email, userData.Password)

	err := res.Scan(&user.Id, &user.Email)
	if err != nil {
		return models.UserDto{}, err
	}

	return user, nil
}

func GetUserByEmail(db *pgxpool.Pool, ctx context.Context, email string) (models.UserDb, error) {
	var user models.UserDb

	res := db.QueryRow(ctx, `
		SELECT id, email, password FROM users WHERE email = $1
	`, email)

	err := res.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return models.UserDb{}, err
	}

	return user, nil
}
