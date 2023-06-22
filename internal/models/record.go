package models

type RecordStatus string

const (
	ProcessingRecordStatus RecordStatus = "processin"
	CompletedRecordStatus  RecordStatus = "completed"
	PublishedRecordStatus  RecordStatus = "published"
)

// RecordCreate model info
//
//	@Description	Record create model
type RecordCreate struct {
	Title     string `json:"title" binding:"required" example:"title"`
	VideoLink string `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
}

// TODO: videolink should be unique?`gorm:"unique"`
type Record struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	VideoLink string
	Status    RecordStatus `gorm:"default:'processing'"`
	Hidden    bool         `gorm:"default:true"`
	UserID    uint
	Article   Article
	Comments  []Comment
}

// RecordDto model info
//
//	@Description	Record dto model
type RecordDto struct {
	ID        uint         `json:"id" binding:"required" example:"1"`
	Title     string       `json:"title" binding:"required" example:"title"`
	VideoLink string       `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
	Status    RecordStatus `json:"status" binding:"required" example:"В обработке"`
	Hidden    bool         `json:"hidden" binding:"required" example:"true"`
}
