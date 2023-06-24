package service

import (
	"errors"
	"profbuh/internal/crud"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateArticleWithRecordID(c *gin.Context, record *models.RecordDto) (models.ArticleDto, error) {
	mlResponse, err := GenerateArticle(c, record)
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

func GetMainArticleForRecord(c *gin.Context, recordID uint) (models.ArticleDto, error) {
	userDb, err := crud.GetUserByEmail(c, c.GetString("x-user-email"))
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.ArticleDto{}, err
	}

	recordDb, err := crud.GetRecordByID(c, recordID)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return models.ArticleDto{}, err
	}

	if recordDb.UserID != userDb.ID && !recordDb.Published {
		return models.ArticleDto{}, err
	}

	article, err := crud.GetMainArticleForRecord(c, recordDb)
	if err != nil {
		logging.Log.Errorf("GetMainArticleForRecord, can't find Article: %v", err)
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticlesForRecord(c *gin.Context, recordID uint, limit int, offset int) ([]models.ArticleDto, error) {
	userDb, err := crud.GetUserByEmail(c, c.GetString("x-user-email"))
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return []models.ArticleDto{}, err
	}

	recordDb, err := crud.GetRecordByID(c, recordID)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return []models.ArticleDto{}, err
	}

	if recordDb.UserID != userDb.ID {
		return []models.ArticleDto{}, errors.New("forbidden")
	}

	articles, err := crud.GetArticlesForRecord(c, recordDb.ID, limit, offset)
	if err != nil {
		logging.Log.Errorf("GetArticlesForRecord, can't find Articles: %v", err)
		return []models.ArticleDto{}, err
	}

	return articles, nil
}

func SetIsMainArticle(c *gin.Context, recordID uint, articleID uint) error {
	userDb, err := crud.GetUserByEmail(c, c.GetString("x-user-email"))
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return err
	}

	recordDb, err := crud.GetRecordByID(c, recordID)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return err
	}

	if recordDb.UserID != userDb.ID {
		return errors.New("forbidden")
	}

	err = crud.SetIsMainArticle(c, recordDb.ID, articleID)
	if err != nil {
		logging.Log.Errorf("SetIsMainArticle, can't set Article as main: %v", err)
		return err
	}

	return nil
}

func CreateAlternativeArticleWithRecordID(c *gin.Context, articleData models.ArticleCreate) (models.ArticleDto, error) {
	userDb, err := crud.GetUserByEmail(c, c.GetString("x-user-email"))
	if err != nil {
		logging.Log.Errorf("GetUserByEmail, can't find email: %v", err)
		return models.ArticleDto{}, err
	}

	recordDb, err := crud.GetRecordByID(c, articleData.RecordID)
	if err != nil {
		logging.Log.Errorf("GetRecordByID, can't find Record: %v", err)
		return models.ArticleDto{}, err
	}

	if recordDb.UserID != userDb.ID {
		return models.ArticleDto{}, errors.New("forbidden")
	}

	article, err := crud.CreateAlternativeArticleWithRecordID(c, articleData)
	if err != nil {
		logging.Log.Errorf("CreateAlternativeArticleWithRecordID, can't create Article: %v", err)
		return models.ArticleDto{}, err
	}

	return article, nil
}
