package router

import (
	"profbuh/internal/database"
	"profbuh/internal/router/api"
	"profbuh/internal/router/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	R  *gin.Engine
	Db *database.Database
}

func NewRouter(db *database.Database) *Router {
	router := &Router{
		R:  gin.Default(),
		Db: db,
	}

	router.InitRoutes()
	return router
}

func (router *Router) InitRoutes() {
	router.R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	auth.InitUserRoutes(router.R, router.Db)
	api.InitRecordRoutes(router.R, router.Db)
	api.InitArticleRoutes(router.R, router.Db)
	api.InitCommentRoutes(router.R, router.Db)

	router.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
