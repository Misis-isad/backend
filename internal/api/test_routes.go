package api

import (
	"net/http"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TestMiddleware godoc
//
//	@Summary		Test middleware
//	@Description	Test middleware
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	models.User	"Current user"
//	@Failure		401	{string}	string		"Unauthorized"
//	@Router			/api/v1/test [get]
func TestMiddleware(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	userDb, err := crud.GetUserByEmail(db, c.Request.Context(), c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.UserDto{
		ID:    userDb.ID,
		Email: userDb.Email,
	})
}
