package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRecord(c *gin.Context, recordData models.RecordCreate, userDb models.User) (models.Record, error) {
	db := c.MustGet("db").(*gorm.DB)

	recordDb := models.Record{
		VideoLink: recordData.VideoLink,
		UserID:    userDb.ID,
		RecordSettings: &models.RecordSettings{
			StartTimecode:    recordData.Settings.StartTimecode,
			EndTimecode:      recordData.Settings.EndTimecode,
			AnnotationLength: recordData.Settings.AnnotationLength,
			ArticleLength:    recordData.Settings.ArticleLength,
			ScreenshotTiming: recordData.Settings.ScreenshotTiming,
		},
	}
	err := db.Model(&models.Record{}).Create(&recordDb).Error
	if err != nil {
		return models.Record{}, err
	}

	err = db.Model(&userDb).Association("Records").Append(&recordDb)
	if err != nil {
		return models.Record{}, err
	}

	return recordDb, nil
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

func SetPublishedStatus(c *gin.Context, recordID uint, userDb models.User, published bool) error {
	db := c.MustGet("db").(*gorm.DB)

	var recordDb models.Record
	err := db.Model(&models.Record{}).Where("id = ?", recordID).Where("user_id = ?", userDb.ID).Where("status = ?", models.RecordStatusReady).First(&recordDb).Error
	if err != nil {
		return err
	}

	err = db.Model(&recordDb).Update("published", published).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPublishedRecords(c *gin.Context, limit int, offset int) ([]models.RecordDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var records []models.RecordDto
	err := db.Model(&models.Record{}).Where("published = ?", true).Limit(limit).Offset(offset).Find(&records).Error
	if err != nil {
		return []models.RecordDto{}, err
	}

	return records, nil
}

func DeleteRecord(c *gin.Context, recordID uint, userDb models.User) error {
	db := c.MustGet("db").(*gorm.DB)

	var recordDb models.Record
	err := db.Model(&models.Record{}).Where("id = ?", recordID).Where("user_id = ?", userDb.ID).First(&recordDb).Error
	if err != nil {
		return err
	}

	err = db.Delete(&recordDb).Error
	if err != nil {
		return err
	}

	return nil
}

func GetRecordByArticleID(c *gin.Context, articleID uint) (models.RecordDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var articleDb models.Article
	err := db.Model(&models.Article{}).Where("id = ?", articleID).First(&articleDb).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	var recordDb models.Record
	err = db.Model(&models.Record{}).Where("id = ?", articleDb.RecordID).First(&recordDb).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	return recordDb.ToDto(), nil
}
