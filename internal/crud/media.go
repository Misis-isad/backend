package crud

import (
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMedias(c *gin.Context, urls []string) ([]models.Media, error) {
	db := c.MustGet("db").(*gorm.DB)

	mediasDb := []models.Media{}

	for _, media := range urls {
		mediaDb := models.Media{
			Link: media,
		}

		mediasDb = append(mediasDb, mediaDb)
	}
	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&mediasDb).Error
	if err != nil {
		return []models.Media{}, err
	}

	return mediasDb, nil
}
