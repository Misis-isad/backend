package api

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
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
//	@Failure		401	{string}	string				"Unauthorized"
//	@Router			/record/create [post]
func (api *ApiClient) CreateRecord(c *gin.Context) {
	var recordData models.RecordCreate

	if err := c.ShouldBindJSON(&recordData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	record, err := service.CreateRecord(api.db.Pool, c, recordData, c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetRecordById godoc
//
//	@Summary		Get record by id
//	@Description	Get record by id
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@Param			record_id	path	int	true	"Record id"
//	@Security		Bearer
//	@Success		200	{object}	models.RecordDto	"Record"
//	@Failure		400	{string}	string				"Bad request"
//	@Failure		401	{string}	string				"Unauthorized"
//	@Router			/record/{record_id} [get]
func (api *ApiClient) GetRecordById(c *gin.Context) {
	recordId, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	record, err := service.GetRecordById(api.db.Pool, c, recordId, c.GetString("x-user-email"))
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
//	@Failure		401	{string}	string				"Unauthorized"
//	@Router			/record/all [get]
func (api *ApiClient) GetRecordsByUser(c *gin.Context) {
	records, err := service.GetRecordsByUser(api.db.Pool, c, c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, records)
}
