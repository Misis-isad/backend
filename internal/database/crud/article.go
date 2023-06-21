package crud

import (
	"context"
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateArticleWithRecordID(db *gorm.DB, c context.Context, articleData models.ArticleCreate, recordID uint) (models.ArticleDto, error) {
	var article models.ArticleDto

	err := db.Model(&models.Article{}).Create(&models.Article{
		RecordID: recordID,
		Body:     articleData.Body,
	}).Preload("Records").Scan(&article).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	err = db.Model(&models.Record{}).Where("id = ?", recordID).Updates(models.Record{Status: models.CompletedRecordStatus}).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticleForRecord(db *gorm.DB, c context.Context, recordID uint) (models.ArticleDto, error) {
	var article models.ArticleDto

	err := db.Model(&models.Article{}).Where("record_id = ?", recordID).First(&article).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}
