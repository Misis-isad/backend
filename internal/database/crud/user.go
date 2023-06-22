package crud

import (
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, userData models.UserCreate) (models.UserDto, error) {
	var user models.UserDto
	err := db.Model(&models.User{}).Create(&models.User{
		Email:    userData.Email,
		Password: userData.Password,
		Fio:      userData.Fio,
	}).Scan(&user).Error

	if err != nil {
		return models.UserDto{}, err
	}

	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	err := db.Model(&models.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
