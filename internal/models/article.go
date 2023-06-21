package models

import "time"

// ArticleCreate model info
//
//	@Description	Article create model
type ArticleCreate struct {
	Body string `json:"body" binding:"required" example:"{html page}" format:"html"`
}

// ArticleDb model info
//
//	@Description	Article db model
type ArticleDb struct {
	Id        int
	Body      string
	RecordId  int
	CreatedAt time.Time
}

// ArticleDto model info
//
//	@Description	Article dto model
type ArticleDto struct {
	Id       int    `json:"id" binding:"required" example:"1"`
	Body     string `json:"body" binding:"required" example:"{html page}" format:"html"`
	RecordId int    `json:"record_id" binding:"required" example:"1"`
}
