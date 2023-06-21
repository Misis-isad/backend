package api

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
//
//	@Summary		Create article
//	@Description	Create article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			record_id	query	int						true	"Record id"
//	@Param			article		body	models.ArticleCreate	true	"Article create info"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto
//	@Failure		400	{object}	string	"Bad request"
//	@Failure		401	{object}	string	"Unauthorized"
//	@Router			/article/create [post]
func (api *ApiClient) CreateArticleWithRecordId(c *gin.Context) {
	var articleData models.ArticleCreate

	if err := c.ShouldBindJSON(&articleData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	recordId, err := strconv.Atoi(c.Query("record_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	article, err := service.CreateArticleWithRecordId(api.db.Pool, c, articleData, recordId)
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
//	@Param			record_id	path	int	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.ArticleDto	"Article"
//	@Failure		400	{string}	string				"Bad request"
//	@Failure		401	{object}	string				"Unauthorized"
//	@Router			/article/{record_id} [get]
func (api *ApiClient) GetArticleForRecord(c *gin.Context) {
	recordId, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	article, err := service.GetArticleForRecord(api.db.Pool, c, recordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}
