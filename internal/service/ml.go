package service

import (
	"os"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

func GenerateArticle(c *gin.Context, record *models.RecordDto) (models.MlResponse, error) {
	// запрос к МЛ для получения статьи
	// получаем body, title, previewPicture, urls

	// FIXME: починить запрос к МЛ
	// r, err := http.Get("http://larek.itatmisis.ru:10000/static/file/eac0a7ec83537763d3ba7671828d0989")
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }

	// articleBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }

	articleBody, err := os.ReadFile("../../test.html")
	if err != nil {
		return models.MlResponse{}, err
	}

	mlResponse := models.MlResponse{
		Body:           string(articleBody),
		Title:          "record title",
		PreviewPicture: "preview_url",
	}

	return mlResponse, nil
}
