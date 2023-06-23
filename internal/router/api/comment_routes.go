package api

import (
	"net/http"
	"profbuh/internal/middlewares"
	"profbuh/internal/models"
	"profbuh/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCommentRoutes(r *gin.Engine) {
	router := r.Group("/api/v1/comment")
	router.Use(middlewares.JwtAuth())
	{
		router.POST("/create", CreateCommentWithRecordID)
		router.GET("/:record_id", GetCommentsForRecord)
		router.GET("/author/:user_id", GetCommentsForUser)
	}
}

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
