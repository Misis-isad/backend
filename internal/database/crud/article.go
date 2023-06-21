package crud

import (
	"context"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateArticleWithRecordId(db *pgxpool.Pool, c context.Context, articleData models.ArticleCreate, recordId int) (models.ArticleDto, error) {
	var article models.ArticleDto

	res := db.QueryRow(c, `
		INSERT INTO articles (body, record_id) VALUES ($1, $2) RETURNING id, body, record_id
	`, articleData.Body, recordId)

	err := res.Scan(&article.Id, &article.Body, &article.RecordId)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, err
}

func GetArticleForRecord(db *pgxpool.Pool, c context.Context, recordId int) (models.ArticleDto, error) {
	var article models.ArticleDb

	res := db.QueryRow(c, `
		SELECT id, body, record_id, created_at FROM articles WHERE record_id = $1
	`, recordId)

	err := res.Scan(&article.Id, &article.Body, &article.RecordId, &article.CreatedAt)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return models.ArticleDto{
		Id:       article.Id,
		Body:     article.Body,
		RecordId: article.RecordId,
	}, nil
}
