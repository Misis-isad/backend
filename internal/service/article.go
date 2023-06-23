package service

import (
	"profbuh/internal/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateArticleWithRecordID(c *gin.Context, record *models.RecordDto) (models.ArticleDto, error) {
	mlResponse, err := GenerateArticle(c, record.VideoLink)
	if err != nil {
		logging.Log.Errorf("GenerateArticle, can't get Article body from ML: %v", err)
		return models.ArticleDto{}, err
	}

	record.Title = mlResponse.Title
	record.PreviewPicture = mlResponse.PreviewPicture

	article, err := crud.CreateArticleWithRecordID(c, record.ID, mlResponse)
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
