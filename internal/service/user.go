package service

import (
	"profbuh/internal/database/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context, userData models.UserCreate) (models.UserDto, error) {
	var err error
	userData.Password, err = HashPassword(userData.Password)
	if err != nil {
		logging.Log.Errorf("CreateUser, can't hash password: %v", err)
		return models.UserDto{}, err
	}

	user, err := crud.CreateUser(c, userData)
	if err != nil {
		logging.Log.Errorf("CreateUser, can't add User to db: %v", err)
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
