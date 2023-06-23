package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRecord(c *gin.Context, recordData models.RecordCreate, userDb models.User) (models.RecordDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	recordDb := models.Record{
		Title:     recordData.Title,
		VideoLink: recordData.VideoLink,
		UserID:    userDb.ID,
	}
	err := db.Model(&models.Record{}).Create(&recordDb).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	return models.RecordDto{
		ID:        recordDb.ID,
		Title:     recordDb.Title,
		VideoLink: recordDb.VideoLink,
		Status:    recordDb.Status,
		Hidden:    recordDb.Hidden,
	}, nil
}

func GetRecordByID(c *gin.Context, recordID uint) (models.Record, error) {
	db := c.MustGet("db").(*gorm.DB)

	var recordDb models.Record
	err := db.Model(&models.Record{}).Where("id = ?", recordID).First(&recordDb).Error
	if err != nil {
		return models.Record{}, err
	}

	return recordDb, nil
}

func GetRecordsForUser(c *gin.Context, userDb models.User, limit int, offset int) ([]models.RecordDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var records []models.RecordDto
	err := db.Model(&models.Record{}).Where("user_id = ?", userDb.ID).Limit(limit).Offset(offset).Find(&records).Error
	if err != nil {
		return []models.RecordDto{}, err
	}

	return records, nil
}

func PublishRecord(c *gin.Context, recordID uint, userDb models.User) (models.RecordDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var recordDb models.Record
	err := db.Model(&models.Record{}).Where("id = ?", recordID).Where("user_id = ?", userDb.ID).First(&recordDb).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	recordDb.Status = models.PublishedRecordStatus
	recordDb.Hidden = false
	err = db.Save(&recordDb).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	return models.RecordDto{
		ID:        recordDb.ID,
		Title:     recordDb.Title,
		VideoLink: recordDb.VideoLink,
		Status:    recordDb.Status,
		Hidden:    recordDb.Hidden,
	}, nil
}
