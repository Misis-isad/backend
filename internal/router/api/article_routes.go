package api

import (
	"net/http"
	"profbuh/internal/middlewares"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitArticleRoutes(r *gin.Engine) {
	router := r.Group("/api/v1/article")
	router.Use(middlewares.JwtAuth())
	{
		router.POST("/create", CreateArticleWithRecordID)
		router.GET("/:record_id", GetArticleByRecordID)
	}
}

// CreateArticle godoc
//
//	@Summary		Create article
//	@Description	Create article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article	body	models.ArticleCreate	true	"Article create info"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto
//	@Failure		400	{object}	string	"Bad request"
//	@Failure		422	{object}	string	"Unprocessable entity"
//	@Router			/api/v1/article/create [post]
func CreateArticleWithRecordID(c *gin.Context) {
	var articleData models.ArticleCreate

	if err := c.ShouldBindJSON(&articleData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	article, err := service.CreateArticleWithRecordID(c, articleData)
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
//	@Failure		404	{object}	string				"Article not found"
//	@Router			/api/v1/article/{record_id} [get]
func GetArticleByRecordID(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	article, err := service.GetArticleForRecord(c, uint(recordID))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}
