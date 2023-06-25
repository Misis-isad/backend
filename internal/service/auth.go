package service

import (
	"errors"
	"profbuh/internal/config"
	"profbuh/internal/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccessToken(email string) (string, error) {
	key := []byte(config.Cfg.JwtSecret)

	claims := Claims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   email,
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

func AuthenticateUser(c *gin.Context, userData models.UserLogin) (string, error) {
	userDb, err := crud.GetUserByEmail(c, userData.Email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return "", err
	}

	if err := VerifyPassword(userData.Password, userDb.Password); err != nil {
		logging.Log.Errorf("VerifyPassword, can't verify password: %v", err)
		return "", err
	}

	token, err := CreateAccessToken(userDb.Email)
	if err != nil {
		logging.Log.Errorf("CreateAccessToken, can't create access token: %v", err)
		return "", err
	}

	return token, nil
}
