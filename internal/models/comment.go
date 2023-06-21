package models

import "time"

// CommentCreate model info
//
//	@Description	Comment create model
type CommentCreate struct {
	Comment string `json:"comment" binding:"required" example:"sample comment"`
}

// CommentDb model info
//
//	@Description	Comment db model
type CommentDb struct {
	ID        int
	Comment   string
	RecordID  int
	AuthorID  int
	CreatedAt time.Time
}

// CommentDto model info
//
//	@Description	Comment dto model
type CommentDto struct {
	ID        int       `json:"id" binding:"required" example:"1"`
	Comment   string    `json:"comment" binding:"required" example:"sample comment"`
	CreatedAt time.Time `json:"created_at" binding:"required" example:"2021-01-01T00:00:00Z"`
}