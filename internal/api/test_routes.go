package api

import (
	"net/http"
	"profbuh/internal/database/crud"
	"profbuh/internal/models"

	"github.com/gin-gonic/gin"
)

// TestMiddleware godoc
//
//	@Summary		Test middleware
//	@Description	Test middleware
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	models.UserDb	"Current user"
//	@Failure		401	{string}	string			"Unauthorized"
//	@Router			/api/v1/test [get]
func (api *ApiClient) TestMiddleware(c *gin.Context) {
	userDb, err := crud.GetUserByEmail(api.db.Pool, c.Request.Context(), c.GetString("x-user-email"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.UserDto{
		Id:    userDb.Id,
		Email: userDb.Email,
	})
}
