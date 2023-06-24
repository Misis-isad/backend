package service

import (
	"context"
	"os"
	"profbuh/internal/models"
	"time"

	"gorm.io/gorm"
)

func GenerateArticle(ctx context.Context, db *gorm.DB, record models.Record) (models.MlResponse, error) {
	// json, err := json.Marshal(record)
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }

	// logging.Log.Debugf("Sending POST request with %v to larek", string(json))
	// request, err := http.NewRequestWithContext(c, "POST", "http://larek.itatmisis.ru:10001/generate_article", bytes.NewBuffer(json))
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }
	// request.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(request)
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	errDesc, _ := io.ReadAll(resp.Body)
	// 	logging.Log.Debug(string(errDesc))
	// 	return models.MlResponse{}, fmt.Errorf("%d", resp.StatusCode)
	// }
	// defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return models.MlResponse{}, err
	// }
	// logging.Log.Debugf("Received: %v", string(data))

	articleBody, err := os.ReadFile("test.html")
	if err != nil {
		return models.MlResponse{}, err
	}

	time.Sleep(25 * time.Second)

	mlResponse := models.MlResponse{
		Body:           string(articleBody),
		Title:          "record title",
		PreviewPicture: "preview_url",
	}

	return mlResponse, nil
}
