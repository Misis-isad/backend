package api

import (
	"net/http"
	"profbuh/internal/database"
	"profbuh/internal/models"
	"profbuh/internal/service"

	"github.com/gin-gonic/gin"
)

type ApiClient struct {
	db *database.Db
}

func NewApiClient(db *database.Db) *ApiClient {
	return &ApiClient{
		db: db,
	}
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserCreate	true	"User create info"
//	@Success		201		{object}	models.UserDto		"Created user"
//	@Failure		400		{string}	string				"Bad request"
//	@Router			/auth/user/create [post]
func (api *ApiClient) CreateUser(c *gin.Context) {
	var userData models.UserCreate

	err := c.ShouldBindJSON(&userData) // body -> json
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = service.CreateUser(api.db.Pool, c.Request.Context(), userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, _ := service.AuthenticateUser(api.db.Pool, c.Request.Context(), models.UserLogin(userData))

	c.JSON(http.StatusCreated, gin.H{"token": token, "token_type": "bearer"})
}

// LoginUser godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserLogin	true	"User login info"
//	@Success		200		{object}	string				"Token"
//	@Failure		401		{string}	string				"Unauthorized"
//	@Router			/auth/user/login [post]
func (api *ApiClient) LoginUser(c *gin.Context) {
	var userData models.UserLogin

	err := c.ShouldBindJSON(&userData) // body -> json
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.AuthenticateUser(api.db.Pool, c.Request.Context(), userData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "token_type": "bearer"})
}
