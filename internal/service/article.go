package service

import (
	"profbuh/internal/database/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateArticleWithRecordID(c *gin.Context, articleData models.ArticleCreate) (models.ArticleDto, error) {
	userDb, err := crud.GetUserByEmail(c, c.GetString("x-user-email"))
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.ArticleDto{}, err
	}

	article, err := crud.CreateArticleWithRecordID(c, articleData, userDb)
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
