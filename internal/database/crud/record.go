package crud

import (
	"context"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRecord(db *pgxpool.Pool, c context.Context, recordData models.RecordCreate, userDb models.UserDb) (models.RecordDto, error) {
	var record models.RecordDto

	res := db.QueryRow(c, `
		INSERT INTO records (title, video_link, author_id) VALUES ($1, $2, $3) RETURNING id, title, video_link, status, visibility, author_id
	`, recordData.Title, recordData.VideoLink, userDb.Id)

	err := res.Scan(&record.Id, &record.Title, &record.VideoLink, &record.Status, &record.Visibility, &record.AuthorId)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, err
}

func GetRecordById(db *pgxpool.Pool, c context.Context, recordId int) (models.RecordDto, error) {
	var record models.RecordDto

	res := db.QueryRow(c, `
		SELECT id, title, video_link, status, visibility, author_id FROM records WHERE id = $1
	`, recordId)

	err := res.Scan(&record.Id, &record.Title, &record.VideoLink, &record.Status, &record.Visibility, &record.AuthorId)
	if err != nil {
		return models.RecordDto{}, err
	}

	return record, err
}

func GetRecordsByUser(db *pgxpool.Pool, c context.Context, userDb models.UserDb) ([]models.RecordDto, error) {
	var records []models.RecordDto

	rows, err := db.Query(c, `
		SELECT id, title, video_link, status, visibility, author_id FROM records WHERE author_id = $1
	`, userDb.Id)
	if err != nil {
		logging.Log.Debug(err)
		return []models.RecordDto{}, err
	}

	for rows.Next() {
		var record models.RecordDto
		err := rows.Scan(&record.Id, &record.Title, &record.VideoLink, &record.Status, &record.Visibility, &record.AuthorId)
		if err != nil {
			return []models.RecordDto{}, err
		}
		records = append(records, record)
	}

	return records, err
}
