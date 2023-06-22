package api

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCommentWithRecordID(c *gin.Context) {
	var commentData models.CommentCreate
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&commentData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	recordID := c.GetUint("record_id")

	comment, err := service.CreateCommentWithRecordID(db, recordID, commentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func GetCommentsForRecord(c *gin.Context) {

}

func GetCommentsForUser(c *gin.Context) {

}
