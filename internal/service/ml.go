package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"profbuh/internal/config"
	"profbuh/internal/logging"
	"profbuh/internal/models"
	"time"

	"gorm.io/gorm"
)

func GenerateArticle(ctx context.Context, db *gorm.DB, record models.Record) (models.MlResponse, error) {
	var mlResponse models.MlResponse

	if config.Cfg.MlService == "true" {
		jsonRecord, err := json.Marshal(record)
		if err != nil {
			return models.MlResponse{}, err
		}

		logging.Log.Debugf("Sending POST request with %v to larek", string(jsonRecord))

		request, err := http.Post("http://larek.itatmisis.ru:10001/generate_article", "application/json", bytes.NewBuffer(jsonRecord))
		if err != nil {
			return models.MlResponse{}, err
		}
		logging.Log.Debugf("Request: %v", request)

		if request.StatusCode != http.StatusOK {
			return models.MlResponse{}, fmt.Errorf("%d", request.StatusCode)
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			return models.MlResponse{}, err
		}
		logging.Log.Debugf("Received: %v", string(body))

		err = json.Unmarshal(body, &mlResponse)
		if err != nil {
			return models.MlResponse{}, err
		}
		logging.Log.Debugf("Unmarshalled: %v", mlResponse)
	} else {

		time.Sleep(15 * time.Second)

		var file, title string
		switch record.VideoLink {
		case "https://www.youtube.com/watch?v=PO2QOECuZT4":
			file = "short"
			title = "Новые требования к учёту выплат в 1С"
		case "https://www.youtube.com/watch?v=rVUHUgEO6qE":
			file = "medium"
			title = "Возврат товаров продавцу на общей системе налогообложения от покупателя на УСН: особенности учета"
		case "https://www.youtube.com/watch?v=J4xHVccHTHk":
			file = "large"
			title = "Руководство по единому налоговому счету: уведомления, платежи и формирование проводок"
		}

		articleBody, err := os.ReadFile("static/" + file + ".html")
		if err != nil {
			return models.MlResponse{}, err
		}

		urls := make(map[string]string, 1)
		urls["https://picsum.photos/200/300"] = "https://picsum.photos/200/300"

		mlResponse = models.MlResponse{
			Body:   string(articleBody),
			Title:  title,
			Images: urls,
		}
	}

	return mlResponse, nil
}
