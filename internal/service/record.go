package service

import (
	"errors"
	"profbuh/internal/database/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context, recordData models.RecordCreate, email string) (models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.RecordDto{}, err
	}

	record, err := crud.CreateRecord(c, recordData, userDb)
	if err != nil {
		logging.Log.Errorf("CreateRecord, can't add Record to db: %v", err)
		return models.RecordDto{}, err
	}

	return record, nil
}

func GetRecordByID(c *gin.Context, recordID uint, email string) (models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.RecordDto{}, err
	}

	recordDb, err := crud.GetRecordByID(c, recordID)
	logging.Log.Debugf("recordDb: %v", recordDb)
	logging.Log.Debugf("err: %v", err)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return models.RecordDto{}, err
	}

	if recordDb.UserID != userDb.ID && recordDb.Hidden {
		return models.RecordDto{}, errors.New("hidden record")
	}

	return models.RecordDto{
		ID:        recordDb.ID,
		Title:     recordDb.Title,
		VideoLink: recordDb.VideoLink,
		Status:    recordDb.Status,
		Hidden:    recordDb.Hidden,
	}, nil
}

func GetRecordsForUser(c *gin.Context, email string, limit int, offset int) ([]models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return []models.RecordDto{}, err
	}

	records, err := crud.GetRecordsForUser(c, userDb, limit, offset)
	if err != nil {
		logging.Log.Errorf("GetRecordsForUser, can't find Records: %v", err)
		return []models.RecordDto{}, err
	}

	return records, nil
}

func PublishRecord(c *gin.Context, recordID uint, email string) (models.RecordDto, error) {
	user, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.RecordDto{}, err
	}

	record, err := crud.PublishRecord(c, recordID, user)
	if err != nil {
		logging.Log.Errorf("PublishRecord, can't publish Record: %v", err)
		return models.RecordDto{}, err
	}

	return record, nil
}
