package service

import (
	"profbuh/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateArticle(c *gin.Context, videoLink string) (models.MlResponse, error) {
	// запрос к МЛ для получения статьи
	// получаем body, title, previewPicture, urls

	time.Sleep(3 * time.Second)

	// urls := []string{"https://i.pinimg.com/736x/f4/d2/96/f4d2961b652880be432fb9580891ed62.jpg"}

	mlResponse := models.MlResponse{
		Body:           "article body " + time.Now().String(),
		Title:          "record title " + time.Now().String(),
		PreviewPicture: "preview_url",
		// MediaLinks:     urls,
	}

	return mlResponse, nil
}
