package service

import (
	"context"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateArticleWithRecordID(db *gorm.DB, c context.Context, articleData models.ArticleCreate, recordID uint) (models.ArticleDto, error) {
	article, err := crud.CreateArticleWithRecordID(db, c, articleData, recordID)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticleForRecord(db *gorm.DB, c context.Context, recordID uint) (models.ArticleDto, error) {
	article, err := crud.GetArticleForRecord(db, c, recordID)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}
