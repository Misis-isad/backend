package middlewares

import (
	"context"
	"profbuh/internal/database"
	"time"

	"github.com/gin-gonic/gin"
)

func DbSession(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeoutContext, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		c.Set("db", db.Db.WithContext(timeoutContext))
		c.Next()
	}
}
