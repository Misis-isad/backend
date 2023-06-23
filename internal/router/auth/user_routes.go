package auth

import (
	"net/http"
	"profbuh/internal/models"
	"profbuh/internal/service"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	router := r.Group("/auth")
	{
		router.POST("/user/create", CreateUser)
		router.POST("/user/login", LoginUser)
	}
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserCreate		true	"User create info"
//	@Success		201		{object}	models.TokenResponse	"Created token for user"
//	@Failure		400		{string}	string					"Bad request"
//	@Failure		422		{string}	string					"Unprocessable entity"
//	@Router			/auth/user/create [post]
func CreateUser(c *gin.Context) {
	var userData models.UserCreate

	err := c.ShouldBindJSON(&userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	_, err = service.CreateUser(c, userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := service.CreateAccessToken(userData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.TokenResponse{Token: token, TokenType: "bearer"})
}

// LoginUser godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserLogin		true	"User login info"
//	@Success		200		{object}	models.TokenResponse	"Token"
//	@Failure		401		{string}	string					"Unauthorized"
//	@Failure		422		{string}	string					"Unprocessable entity"
//	@Router			/auth/user/login [post]
func LoginUser(c *gin.Context) {
	var userData models.UserLogin

	err := c.ShouldBindJSON(&userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	token, err := service.AuthenticateUser(c, userData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{Token: token, TokenType: "bearer"})
}
