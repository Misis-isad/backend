package service

import (
	"context"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateRecord(db *gorm.DB, c context.Context, recordData models.RecordCreate, email string) (models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(db, c, email)
	if err != nil {
		return models.RecordDto{}, err
	}

	record, err := crud.CreateRecord(db, c, recordData, userDb)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}

func GetRecordByID(db *gorm.DB, c context.Context, recordID int, email string) (models.RecordDto, error) {
	record, err := crud.GetRecordByID(db, c, recordID)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}

func GetRecordsByUser(db *gorm.DB, c context.Context, email string) ([]models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(db, c, email)
	if err != nil {
		return []models.RecordDto{}, err
	}

	records, err := crud.GetRecordsByUser(db, c, userDb)
	if err != nil {
		return []models.RecordDto{}, err
	}

	return records, nil
}

func PublishRecord(db *gorm.DB, c context.Context, recordID uint, email string) (models.RecordDto, error) {
	user, err := crud.GetUserByEmail(db, c, email)
	if err != nil {
		return models.RecordDto{}, err
	}

	record, err := crud.PublishRecord(db, c, recordID, user)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, nil
}
