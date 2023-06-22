package crud

import (
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateRecord(db *gorm.DB, recordData models.RecordCreate, userDb models.User) (models.RecordDto, error) {
	var record models.RecordDto
	err := db.Model(&models.Record{}).Create(&models.Record{
		UserID:    userDb.ID,
		Title:     recordData.Title,
		VideoLink: recordData.VideoLink,
	}).Scan(&record).Error

	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}

func GetRecordByID(db *gorm.DB, recordID int) (models.RecordDto, error) {
	var record models.RecordDto
	err := db.Model(&models.Record{}).Where("id = ?", recordID).First(&record).Error

	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}

func GetRecordsForUser(db *gorm.DB, userDb models.User) ([]models.RecordDto, error) {
	var records []models.RecordDto

	err := db.Model(&models.Record{}).Where("user_id = ?", userDb.ID).Find(&records).Error
	if err != nil {
		return []models.RecordDto{}, err
	}

	return records, nil
}

func PublishRecord(db *gorm.DB, recordID uint, userDb models.User) (models.RecordDto, error) {
	var record models.RecordDto

	err := db.Model(&models.Record{}).Where("id = ?", recordID).Where("user_id = ?", userDb.ID).Updates(models.Record{
		Status: models.PublishedRecordStatus,
		Hidden: false,
	}).Scan(&record).Error
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}
