package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArticleWithRecordID(c *gin.Context, recordID uint, mlResponse models.MlResponse) (models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	articleDb := models.Article{
		Body:     mlResponse.Body,
		RecordID: recordID,
	}
	err := db.Model(&models.Article{}).Create(&articleDb).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	err = db.Model(&models.Record{}).Where("id = ?", recordID).Updates(&models.Record{
		Title:          mlResponse.Title,
		PreviewPicture: mlResponse.PreviewPicture,
	}).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	// mediasDb, err := CreateMedias(c, urls)
	// if err != nil {
	// 	return models.ArticleDto{}, err
	// }

	// err = db.Model(&articleDb).Association("MediaLinks").Append(&mediasDb)
	// if err != nil {
	// 	return models.ArticleDto{}, err
	// }

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
