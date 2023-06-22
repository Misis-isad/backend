package service

import (
	"context"
	"errors"
	"profbuh/internal/config"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccessToken(userData models.User) (string, error) {
	key := []byte(config.Cfg.JwtSecret)

	claims := Claims{
		userData.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   userData.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // проверка метода шифрования
			return nil, errors.New("unexpected signing method")
		}
		// email -> userDb
		return []byte(config.Cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(db *gorm.DB, c context.Context, userData models.UserLogin) (string, error) {
	userDb, err := crud.GetUserByEmail(db, userData.Email)
	if err != nil {
		return "", err
	}

	if ok := VerifyPassword(userData.Password, userDb.Password); ok != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := CreateAccessToken(userDb)
	if err != nil {
		return "", err
	}

	return token, nil
}
