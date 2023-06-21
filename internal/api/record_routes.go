package api

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
//	@Router			/record/create [post]
func CreateRecord(c *gin.Context) {
	var recordData models.RecordCreate
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&recordData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	record, err := service.CreateRecord(db, c, recordData, c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, record)
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
//	@Failure		400	{string}	string				"Bad request"
//	@Router			/record/{record_id} [get]
func GetRecordByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	recordID, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	record, err := service.GetRecordByID(db, c, recordID, c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
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
//	@Security		Bearer
//	@Success		200	{array}		models.RecordDto	"User's records"
//	@Failure		400	{string}	string				"Bad request"
//	@Router			/record/all [get]
func GetRecordsByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	records, err := service.GetRecordsByUser(db, c, c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, records)
}

// PublishRecord godoc
//
//	@Summary		Publish record
//	@Description	Publish record
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	uint	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.RecordDto
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/record/{record_id}/publish [post]
func PublishRecord(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	recordID, err := strconv.ParseUint(c.Param("record_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	record, err := service.PublishRecord(db, c, uint(recordID), c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, record)
}
