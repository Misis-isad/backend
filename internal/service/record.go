package service

import (
	"errors"
	"profbuh/internal/crud"
	"profbuh/internal/database"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context, recordData models.RecordCreate, email string, db *database.Database) (models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.RecordDto{}, err
	}

	recordDb, err := crud.CreateRecord(c, recordData, userDb)
	if err != nil {
		logging.Log.Errorf("CreateRecord, can't add Record to db: %v", err)
		return models.RecordDto{}, err
	}

	logging.Log.Debug("Starting background ml create article")
	go BackgroundMlCreateArticle(recordDb, db)

	return recordDb.ToDto(), nil
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

	if recordDb.UserID != userDb.ID && !recordDb.Published {
		return models.RecordDto{}, errors.New("unpublished record")
	}

	return recordDb.ToDto(), nil
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

func SetPublishedStatus(c *gin.Context, recordID uint, email string, published bool) error {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return err
	}

	err = crud.SetPublishedStatus(c, recordID, userDb, published)
	if err != nil {
		logging.Log.Errorf("SetPublishedStatus, can't publish Record: %v", err)
		return err
	}

	return nil
}

func GetPublishedRecords(c *gin.Context, limit int, offset int) ([]models.RecordDto, error) {
	records, err := crud.GetPublishedRecords(c, limit, offset)
	if err != nil {
		logging.Log.Errorf("GetPublishedRecords, can't find Records: %v", err)
		return []models.RecordDto{}, err
	}

	return records, nil
}

func DeleteRecord(c *gin.Context, recordID uint, email string) error {
	userDb, err := crud.GetUserByEmail(c, email)
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return err
	}

	err = crud.DeleteRecord(c, recordID, userDb)
	if err != nil {
		logging.Log.Errorf("DeleteRecord, can't delete Record: %v", err)
		return err
	}

	return nil
}
