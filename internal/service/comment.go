package service

import (
	"profbuh/internal/crud"
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateCommentWithRecordID(db *gorm.DB, recordID uint, commentData models.CommentCreate) (models.CommentDto, error) {
	comment, err := crud.CreateCommentWithRecordID(db, recordID, commentData)
	if err != nil {
		return models.CommentDto{}, err
	}

	return comment, nil
}
