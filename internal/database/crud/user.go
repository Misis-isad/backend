package crud

import (
	"context"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(db *pgxpool.Pool, c context.Context, userData models.UserCreate) (models.UserDto, error) {
	var user models.UserDto

	res := db.QueryRow(c, `
		INSERT INTO users (email, password, fio) VALUES ($1, $2, $3) RETURNING id, email, fio
	`, userData.Email, userData.Password, userData.Fio)

	err := res.Scan(&user.Id, &user.Email, &user.Fio)
	if err != nil {
		logging.Log.Debug(err.Error())
		return models.UserDto{}, err
	}

	return user, err
}

func GetUserByEmail(db *pgxpool.Pool, c context.Context, email string) (models.UserDb, error) {
	var user models.UserDb

	res := db.QueryRow(c, `
		SELECT id, email, password, fio FROM users WHERE email = $1
	`, email)

	err := res.Scan(&user.Id, &user.Email, &user.Password, &user.Fio)
	if err != nil {
		logging.Log.Debug(err.Error())
		return models.UserDb{}, err
	}

	return user, err
}
