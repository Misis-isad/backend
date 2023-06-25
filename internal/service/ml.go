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

	"gorm.io/gorm"
)

func GenerateArticle(ctx context.Context, db *gorm.DB, record models.Record) (models.MlResponse, error) {
	var mlResponse models.MlResponse

	var file, title string
	urls := make(map[string]string, 0)
	switch record.VideoLink {
	// case "https://www.youtube.com/watch?v=PO2QOECuZT4":
	// 	file = "short"
	// 	title = "Новые требования к учёту выплат в 1С"
	// 	urls["00:00:15"] = "http://larek.itatmisis.ru:10000/static/file/caaa6c2847df4b7c80169ce774139f1b.jpg"
	// 	urls["00:01:22"] = "http://larek.itatmisis.ru:10000/static/file/83a87596ef5543c2a4e33e0400eac00a.jpg"
	// 	time.Sleep(20 * time.Second)
	// case "https://www.youtube.com/watch?v=rVUHUgEO6qE":
	// 	file = "medium"
	// 	title = "Возврат товаров продавцу на общей системе налогообложения от покупателя на УСН: особенности учета"
	// 	urls["{int(hours):02d}:{int(minutes):02d}:{int(seconds):02d}.jpg"] = "'http://larek.itatmisis.ru:10000/static/file/6ce380dff40049659e29038e6774f4dd.jpg'"
	// 	time.Sleep(40 * time.Second)
	// case "https://www.youtube.com/watch?v=J4xHVccHTHk":
	// 	file = "large"
	// 	title = "Руководство по единому налоговому счету: уведомления, платежи и формирование проводок"
	// 	urls["{int(hours):02d}:{int(minutes):02d}:{int(seconds):02d}.jpg"] = "'http://larek.itatmisis.ru:10000/static/file/6ce380dff40049659e29038e6774f4dd.jpg'"
	// 	time.Sleep(90 * time.Second)
	default:
		if config.Cfg.MlService == "true" {
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

			// article := mlResponse.Body
			// timecodes := regexp.MustCompile(`\[(.*?)\]`).FindAllString(article, -1)

			// for filename, url := range mlResponse.Images {
			// 	logging.Log.Debugf("Filename: %v, URL: %v", filename, url)
			// 	for _, timecodePair := range timecodes {
			// 		curTimecodes := regexp.MustCompile(`-`).Split(timecodePair, -1)
			// 		start := regexp.MustCompile(":").Split(curTimecodes[0], -1)
			// 		startMinutes, _ := strconv.ParseFloat(start[0], 64)
			// 		startSecondsAndMili, _ := strconv.ParseFloat(start[1], 64)
			// 		startSecondsStr := regexp.MustCompile(".").Split(strconv.FormatFloat(startSecondsAndMili, 'f', -1, 64), -1)[0]
			// 		startSeconds, _ := strconv.ParseFloat(startSecondsStr, 64)
			// 		startDuration := 60*startMinutes + startSeconds

			// 		end := regexp.MustCompile(":").Split(curTimecodes[1], -1)
			// 		endMinutes, _ := strconv.ParseFloat(end[0], 64)
			// 		endSecondsAndMili, _ := strconv.ParseFloat(end[1], 64)
			// 		endSecondsStr := regexp.MustCompile(".").Split(strconv.FormatFloat(endSecondsAndMili, 'f', -1, 64), -1)[0]
			// 		endSeconds, _ := strconv.ParseFloat(endSecondsStr, 64)
			// 		endDuration := 60*endMinutes + endSeconds

			// 		filename := regexp.MustCompile(":").Split(filename, -1)
			// 		filenameHours, _ := strconv.ParseFloat(filename[0], 64)
			// 		filenameMinutes, _ := strconv.ParseFloat(filename[1], 64)
			// 		filenameSeconds, _ := strconv.ParseFloat(filename[2], 64)
			// 		filenameDuration := 60*60*filenameHours + 60*filenameMinutes + filenameSeconds

			// 		logging.Log.Debugf("Start: %v, End: %v, Filename: %v", startDuration, endDuration, filenameDuration)

			// 		if filenameDuration >= startDuration && filenameDuration <= endDuration {
			// 			article = regexp.MustCompile(`\[(.*?)\]`).ReplaceAllString(article, "<img src="+url+">")
			// 		}
			// 	}
			// }

			// logging.Log.Debugf("Article: %v", article)
			// mlResponse.Body = article

			return mlResponse, nil
		} else {
			file = "test"
			title = "ABACABA"
			urls["test"] = "test"
		}
	}

	articleBody, err := os.ReadFile("static/" + file + ".html")
	if err != nil {
		return models.MlResponse{}, err
	}

	article := string(articleBody)

	logging.Log.Debugf("Article: %v", article)

	mlResponse = models.MlResponse{
		Body:   article,
		Title:  title,
		Images: urls,
	}
	return mlResponse, nil
}
