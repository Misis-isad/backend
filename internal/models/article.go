package models

import "time"

// ArticleCreate model info
//
//	@Description	Article create model
type ArticleCreate struct {
	Body     string `json:"body" binding:"required" example:"{html page}" format:"html"`
	RecordID uint   `json:"record_id" binding:"required" example:"1"`
}

type Article struct {
	ID        uint `gorm:"primaryKey"`
	Body      string
	RecordID  uint
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// ArticleDto model info
//
//	@Description	Article dto model
type ArticleDto struct {
	ID        uint      `json:"id" binding:"required" example:"1"`
	Body      string    `json:"body" binding:"required" example:"{html page}" format:"html"`
	CreatedAt time.Time `json:"created_at" binding:"required" example:"2021-01-01T00:00:00Z"`
}
