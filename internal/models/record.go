package models

type RecordStatus string

const (
	ProcessingRecordStatus RecordStatus = "В обработке"
	CompletedRecordStatus  RecordStatus = "Обработано"
	PublicRecordStatus     RecordStatus = "Опубликовано"
)

// RecordCreate model info
//
//	@Description	Record create model
type RecordCreate struct {
	Title     string `json:"title" binding:"required" example:"title"`
	VideoLink string `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
}

// RecordDb model info
//
//	@Description	Record db model
type RecordDb struct {
	Id         int
	Title      string
	VideoLink  string
	Status     RecordStatus
	Visibility bool
	// MainArticleId int
	AuthorId int
}

// RecordDto model info
//
//	@Description	Record dto model
type RecordDto struct {
	Id         int          `json:"id" binding:"required" example:"1"`
	Title      string       `json:"title" binding:"required" example:"title"`
	VideoLink  string       `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
	Status     RecordStatus `json:"status" binding:"required" example:"В обработке"`
	Visibility bool         `json:"visibility" binding:"required" example:"true"`
	// MainArticleId int          `json:"main_article_id" binding:"required" example:"1"`
	AuthorId int `json:"author_id" binding:"required" example:"1"`
}
