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
		router.GET("/:record_id/main", GetMainArticleByRecordID)
		router.GET("/:record_id/all", GetArticlesForRecord)
		router.POST("/:record_id/is_main", SetIsMainArticle)
		router.POST("/alternative", CreateAlternativeArticleWithRecordID)
	}
}

// GetMainArticleByRecordID godoc
//
//	@Summary		Get main article
//	@Description	Get main article for record by record_id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto	"Article"
//	@Failure		404	{object}	string				"Article not found"
//	@Failure		422	{object}	string				"Unprocessable entity"
//	@Router			/api/v1/article/{record_id}/main [get]
func GetMainArticleByRecordID(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	article, err := service.GetMainArticleForRecord(c, uint(recordID))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}

// GetArticlesForRecord godoc
//
//	@Summary		Get articles
//	@Description	Get articles for record by record_id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Param			limit		query	int		false	"Limit"
//	@Param			offset		query	int		false	"Offset"
//	@Security		Bearer
//	@Success		200	{object}	[]models.ArticleDto
//	@Failure		403	{object}	string	"Forbidden"
//	@Failure		404	{object}	string	"Articles not found"
//	@Failure		422	{object}	string	"Unprocessable entity"
//	@Router			/api/v1/article/{record_id}/all [get]
func GetArticlesForRecord(c *gin.Context) {
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	recordId, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	articles, err := service.GetArticlesForRecord(c, uint(recordId), int(limit), int(offset))
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, err.Error())
			return
		} else {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, articles)
}

// SetIsMainArticle godoc
//
//	@Summary		Set is_main
//	@Description	Set is_main for article by record_id and article_id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Param			article_id	query	uint	true	"Article id"
//	@Security		Bearer
//	@Success		204	{object}	string	"Article set as main"
//	@Failure		403	{object}	string	"Forbidden"
//	@Failure		404	{object}	string	"Article not found"
//	@Failure		422	{object}	string	"Unprocessable entity"
//	@Router			/api/v1/article/{record_id}/is_main [post]
func SetIsMainArticle(c *gin.Context) {
	recordId, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	articleId, err := strconv.ParseUint(c.Query("article_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = service.SetIsMainArticle(c, uint(recordId), uint(articleId))
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, err.Error())
			return
		} else {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusNoContent, "Article set as main")
}

// CreateAlternativeArticleWithRecordID godoc
//
//	@Summary		Create alternative article
//	@Description	Create alternative article for record by record_id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article	body	models.ArticleCreate	true	"Article info for create"
//	@Security		Bearer
//	@Success		200	{object}	string	"Article created"
//	@Failure		403	{object}	string	"Forbidden"
//	@Failure		404	{object}	string	"Article not found"
//	@Failure		422	{object}	string	"Unprocessable entity"
//	@Router			/api/v1/article/alternative [post]
func CreateAlternativeArticleWithRecordID(c *gin.Context) {
	var articleData models.ArticleCreate
	if err := c.ShouldBindJSON(&articleData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	article, err := service.CreateAlternativeArticleWithRecordID(c, articleData)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, err.Error())
			return
		} else {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, article)
}
