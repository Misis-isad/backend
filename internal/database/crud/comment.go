package crud

import (
	"profbuh/internal/models"

	"gorm.io/gorm"
)

func CreateCommentWithRecordID(db *gorm.DB, recordID uint, commentData models.CommentCreate) (models.CommentDto, error) {
	var comment models.CommentDto

	err := db.Model(&models.Comment{}).Create(&models.Comment{
		Comment: commentData.Comment,
	}).Preload("Records").Error

	if err != nil {
		return models.CommentDto{}, err
	}

	return comment, nil
}
