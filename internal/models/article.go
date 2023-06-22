package models

// ArticleCreate model info
//
//	@Description	Article create model
type ArticleCreate struct {
	Body string `json:"body" binding:"required" example:"{html page}" format:"html"`
}

type Article struct {
	ID       uint `gorm:"primaryKey"`
	Body     string
	RecordID uint
}

// ArticleDto model info
//
//	@Description	Article dto model
type ArticleDto struct {
	ID   uint   `json:"id" binding:"required" example:"1"`
	Body string `json:"body" binding:"required" example:"{html page}" format:"html"`
}
