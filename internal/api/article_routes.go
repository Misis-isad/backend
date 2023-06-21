package api

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateArticle godoc
//
//	@Summary		Create article
//	@Description	Create article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	query	uint					true	"Record id"
//	@Param			article		body	models.ArticleCreate	true	"Article create info"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto
//	@Failure		400	{object}	string	"Bad request"
//	@Router			/article/create [post]
func CreateArticleWithRecordID(c *gin.Context) {
	var articleData models.ArticleCreate
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&articleData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	recordID, err := strconv.ParseUint(c.Query("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	article, err := service.CreateArticleWithRecordID(db, c, articleData, uint(recordID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}

// GetArticleForRecord godoc
//
//	@Summary		Get article
//	@Description	Get article for record by record_id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto	"Article"
//	@Failure		400	{string}	string				"Bad request"
//	@Router			/article/{record_id} [get]
func GetArticleByRecordID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	article, err := service.GetArticleForRecord(db, c, uint(recordID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}
