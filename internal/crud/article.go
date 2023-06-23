package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArticleWithRecordID(c *gin.Context, articleData models.ArticleCreate) (models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	articleDb := models.Article{
		Body:     articleData.Body,
		RecordID: articleData.RecordID,
	}
	err := db.Model(&models.Article{}).Create(&articleDb).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	return models.ArticleDto{
		ID:   articleDb.ID,
		Body: articleDb.Body,
	}, nil
}

func GetArticleForRecord(c *gin.Context, recordDb models.Record) (models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var article models.ArticleDto
	err := db.Model(&models.Article{}).Where("record_id = ?", recordDb.ID).First(&article).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}
