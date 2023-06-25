package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"profbuh/internal/config"
	"profbuh/internal/logging"
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func GenerateArticle(ctx context.Context, db *gorm.DB, record models.Record) (models.MlResponse, error) {
	var mlResponse models.MlResponse
	jsonRecord, err := json.Marshal(record)
	if err != nil {
		return models.MlResponse{}, err
	}

	logging.Log.Debugf("Sending POST request with %v to larek", string(jsonRecord))

	request, err := http.Post(config.Cfg.LarekUrl+":10001/generate_article", "application/json", bytes.NewBuffer(jsonRecord))
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

	return mlResponse, nil
}
