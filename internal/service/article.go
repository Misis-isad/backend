package service

import (
	"profbuh/internal/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateArticleWithRecordID(c *gin.Context, record models.RecordDto) (models.ArticleDto, error) {
	body, err := GenerateArticle(c, record.VideoLink)
	if err != nil {
		logging.Log.Errorf("GenerateArticle, can't get Article body from ML: %v", err)
	}

	article, err := crud.CreateArticleWithRecordID(c, models.ArticleCreate{
		Body:     body,
		RecordID: record.ID,
	})
	if err != nil {
		logging.Log.Errorf("CreateArticleWithRecordID, can't add Article to db: %v", err)
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticleForRecord(c *gin.Context, recordID uint) (models.ArticleDto, error) {
	recordDb, err := crud.GetRecordByID(c, recordID)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return models.ArticleDto{}, err
	}

	article, err := crud.GetArticleForRecord(c, recordDb)
	if err != nil {
		logging.Log.Errorf("GetArticleForRecord, can't find Article: %v", err)
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GenerateArticle(c *gin.Context, videoLink string) (string, error) {
	// запрос к МЛ для получения статьи

	time.Sleep(3 * time.Second)

	return "article body", nil
}
