package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, userData models.UserCreate) (models.UserDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	userDb := models.User{
		Email:    userData.Email,
		Password: userData.Password,
		Fio:      userData.Fio,
	}
	err := db.Model(&models.User{}).Create(&userDb).Error

	if err != nil {
		return models.UserDto{}, err
	}

	return models.UserDto{
		ID:    userDb.ID,
		Email: userDb.Email,
		Fio:   userDb.Fio,
	}, nil
}

func GetUserByEmail(c *gin.Context, email string) (models.User, error) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	err := db.Model(&models.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
