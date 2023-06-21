package service

import (
	"context"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRecord(db *pgxpool.Pool, c context.Context, recordData models.RecordCreate, email string) (models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(db, c, email)
	if err != nil {
		return models.RecordDto{}, err
	}

	record, err := crud.CreateRecord(db, c, recordData, userDb)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, err
}

func GetRecordById(db *pgxpool.Pool, c context.Context, recordId int, email string) (models.RecordDto, error) {
	// userDb, err := crud.GetUserByEmail(db, c, email)
	// if err != nil {
	// 	return models.RecordDto{}, err
	// }

	record, err := crud.GetRecordById(db, c, recordId)
	if err != nil {
		return models.RecordDto{}, err
	}

	// if userDb.Id != record.AuthorId {
	// 	return models.RecordDto{}, err
	// }

	return record, err
}

func GetRecordsByUser(db *pgxpool.Pool, c context.Context, email string) ([]models.RecordDto, error) {
	userDb, err := crud.GetUserByEmail(db, c, email)
	if err != nil {
		return []models.RecordDto{}, err
	}

	records, err := crud.GetRecordsByUser(db, c, userDb)
	if err != nil {
		return []models.RecordDto{}, err
	}

	return records, err
}
