package service

import (
	"context"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateArticleWithRecordId(db *pgxpool.Pool, c context.Context, articleData models.ArticleCreate, recordId int) (models.ArticleDto, error) {
	article, err := crud.CreateArticleWithRecordId(db, c, articleData, recordId)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticleForRecord(db *pgxpool.Pool, c context.Context, recordId int) (models.ArticleDto, error) {
	article, err := crud.GetArticleForRecord(db, c, recordId)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}
