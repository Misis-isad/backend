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
		IsMain:   true,
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

	// FIXME: не работает
	// err = db.Model(&models.Record{}).Association("Articles").Append(&articleDb)
	// if err != nil {
	// 	return models.ArticleDto{}, err
	// }

	return models.ArticleDto{
		ID:   articleDb.ID,
		Body: articleDb.Body,
	}, nil
}

func GetMainArticleForRecord(c *gin.Context, recordDb models.Record) (models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var article models.ArticleDto
	err := db.Model(&models.Article{}).Where("record_id = ?", recordDb.ID).Where("is_main = ?", true).First(&article).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	return article, nil
}

func GetArticlesForRecord(c *gin.Context, recordID uint, limit int, offset int) ([]models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	var articles []models.ArticleDto
	err := db.Model(&models.Article{}).Where("record_id = ?", recordID).Limit(limit).Offset(offset).Find(&articles).Error
	if err != nil {
		return []models.ArticleDto{}, err
	}

	return articles, nil
}

func SetIsMainArticle(c *gin.Context, recordID uint, articleID uint) error {
	db := c.MustGet("db").(*gorm.DB)

	err := db.Model(&models.Article{}).Where("record_id = ?", recordID).Update("is_main", false).Error
	if err != nil {
		return err
	}

	err = db.Model(&models.Article{}).Where("id = ?", articleID).Updates(&models.Article{
		IsMain: true,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func CreateAlternativeArticleWithRecordID(c *gin.Context, articleData models.ArticleCreate) (models.ArticleDto, error) {
	db := c.MustGet("db").(*gorm.DB)

	articleDb := models.Article{
		Body:     articleData.Body,
		RecordID: articleData.RecordID,
		IsMain:   false,
	}
	err := db.Model(&models.Article{}).Create(&articleDb).Error
	if err != nil {
		return models.ArticleDto{}, err
	}

	// FIXME: не работает
	// err = db.Model(&models.Record{}).Association("Articles").Append(&articleDb)
	// if err != nil {
	// 	return models.ArticleDto{}, err
	// }

	return models.ArticleDto{
		ID:   articleDb.ID,
		Body: articleDb.Body,
	}, nil
}
