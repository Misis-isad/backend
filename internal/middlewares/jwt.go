package middlewares

import (
	"errors"
	"net/http"
	"profbuh/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			token, err := service.VerifyToken(authToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			email := token.Claims.(jwt.MapClaims)["email"]
			c.Set("x-user-email", email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, errors.New("no token"))
			c.Abort()
		}
	}
}
