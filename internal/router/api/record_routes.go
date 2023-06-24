package api

import (
	"net/http"
	"profbuh/internal/database"
	"profbuh/internal/middlewares"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitRecordRoutes(r *gin.Engine, db *database.Database) {
	router := r.Group("/api/v1/record")
	router.Use(middlewares.JwtAuth())
	router.Use(middlewares.DbSession(db, 5))
	{
		router.POST("/create", CreateRecord(db))
		router.GET("/:record_id", GetRecordByID)
		router.GET("/all", GetRecordsForUser)
		router.POST("/:record_id/published_status", SetPublishedStatus)
		router.GET("/published", GetPublishedRecords)
		router.DELETE("/:record_id", DeleteRecord)
		router.GET("/by_article/:article_id", GetRecordByArticleID)
	}
}

// CreateRecord godoc
//
//	@Summary		Create record
//	@Description	Create record
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record	body	models.RecordCreate	true	"Record create info"
//	@Security		Bearer
//	@Success		200	{object}	models.RecordDto	"Created record"
//	@Failure		400	{string}	string				"Bad request"
//	@Failure		422	{string}	string				"Unprocessable entity"
//	@Router			/api/v1/record/create [post]
func CreateRecord(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var recordData models.RecordCreate

		if err := c.ShouldBindJSON(&recordData); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		record, err := service.CreateRecord(c, recordData, c.GetString("x-user-email"), db)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, record)
	}
}

// GetRecordByID godoc
//
//	@Summary		Get record by id
//	@Description	Get record by id
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.RecordDto	"Record"
//	@Failure		403	{string}	string				"Hidden record"
//	@Failure		404	{string}	string				"Record not found"
//	@Failure		422	{string}	string				"Unprocessable entity"
//	@Router			/api/v1/record/{record_id} [get]
func GetRecordByID(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	record, err := service.GetRecordByID(c, uint(recordID), c.GetString("x-user-email"))
	if err != nil {
		if err.Error() == "hidden record" {
			c.JSON(http.StatusForbidden, err.Error())
			return
		} else {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, record)
}

// GetRecordsByUser godoc
//
//	@Summary		Get records by user
//	@Description	Get all records(not articles) for current user
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	int	false	"Limit"		validation:"gte=0,lte=100"	default(10)
//	@Param			offset	query	int	false	"Offset"	validation:"gte=0"			default(0)
//	@Security		Bearer
//	@Success		200	{array}		models.RecordDto	"Records"
//	@Failure		400	{string}	string				"Bad request"
//	@Failure		422	{string}	string				"Unprocessable entity"
//	@Router			/api/v1/record/all [get]
func GetRecordsForUser(c *gin.Context) {
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

	records, err := service.GetRecordsForUser(c, c.GetString("x-user-email"), int(limit), int(offset))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, records)
}

// SetPublishedStatus godoc
//
//	@Summary		Set published status
//	@Description	Set published status
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Param			published	query	bool	true	"Published status"
//	@Security		Bearer
//	@Success		204	{object}	string	"Record published"
//	@Failure		404	{string}	string	"Record not found"
//	@Failure		422	{string}	string	"Unprocessable entity"
//	@Router			/api/v1/record/{record_id}/published_status [post]
func SetPublishedStatus(c *gin.Context) {
	published, err := strconv.ParseBool(c.Query("published"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = service.SetPublishedStatus(c, uint(recordID), c.GetString("x-user-email"), published)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "Record published")
}

// GetPublishedRecords godoc
//
//	@Summary		Get published records
//	@Description	Get all published records(not articles)
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	int	false	"Limit"		validation:"gte=0,lte=100"	default(10)
//	@Param			offset	query	int	false	"Offset"	validation:"gte=0"			default(0)
//	@Security		Bearer
//	@Success		200	{array}		models.RecordDto	"Records"
//	@Failure		400	{string}	string				"Bad request"
//	@Failure		422	{string}	string				"Unprocessable entity"
//	@Router			/api/v1/record/published [get]
func GetPublishedRecords(c *gin.Context) {
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

	records, err := service.GetPublishedRecords(c, int(limit), int(offset))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, records)
}

// DeleteRecord godoc
//
//	@Summary		Delete record
//	@Description	Delete record
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Security		Bearer
//	@Success		204	{string}	string
//	@Failure		404	{string}	string	"Record not found"
//	@Failure		422	{string}	string	"Unprocessable entity"
//	@Router			/api/v1/record/{record_id} [delete]
func DeleteRecord(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = service.DeleteRecord(c, uint(recordID), c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// GetRecordByArticleID godoc
//
//	@Summary		Get record
//	@Description	Get record for article by article_id
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path	uint	true	"Article id"
//	@Security		Bearer
//	@Success		200	{object}	models.RecordDto	"Found record"
//	@Failure		404	{object}	string				"Article not found"
//	@Failure		422	{object}	string				"Unprocessable entity"
//	@Router			/api/v1/record/by_article/{article_id} [get]
func GetRecordByArticleID(c *gin.Context) {
	articleId, err := strconv.ParseUint(c.Param("article_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	record, err := service.GetRecordByArticleID(c, uint(articleId))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, record)
}
