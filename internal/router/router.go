package router

import (
	"profbuh/internal/database"
	"profbuh/internal/middlewares"
	"profbuh/internal/router/api"
	"profbuh/internal/router/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
	Db *database.Database
}

func NewRouter(db *database.Database) *Router {
	router := &Router{
		Engine: gin.Default(),
		Db:     db,
	}

	router.InitRoutes()
	return router
}

func (r *Router) InitRoutes() {
	r.Use(middlewares.DbSession(r.Db))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	auth.InitUserRoutes(r.Engine)
	api.InitRecordRoutes(r.Engine)
	api.InitArticleRoutes(r.Engine)
	api.InitCommentRoutes(r.Engine)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
