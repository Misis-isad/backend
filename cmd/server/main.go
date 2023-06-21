package main

import (
	"fmt"
	"os"
	"profbuh/internal/api"
	"profbuh/internal/config"
	"profbuh/internal/database"
	"profbuh/internal/logging"
	"profbuh/middlewares"

	_ "profbuh/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Profbuh API
//	@description	This is a sample server for Profbuh API.
//
//	@host			localhost:8000
func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Can't init log: %v", err)
		os.Exit(1)
	}

	err = logging.InitLog()
	if err != nil {
		fmt.Printf("Can't init log: %v", err)
		os.Exit(1)
	}

	logging.Log.Info("Starting server")

	db, err := database.InitDb(&config.Cfg)
	if err != nil {
		logging.Log.Fatalf("Can't init db: %v", err)
	}

	logging.Log.Debug("Db connected")

	gin.SetMode(config.Cfg.GinMode)
	r := gin.Default()
	r.Use(middlewares.DbSession(db))

	router_auth := r.Group("/auth")
	{
		router_auth.POST("/user/create", api.CreateUser)
		router_auth.POST("/user/login", api.LoginUser)
	}

	router_api := r.Group("/api/v1")
	router_api.Use(middlewares.JwtAuth())
	{
		router_api.GET("/test", api.TestMiddleware)

		router_api.POST("/article/create", api.CreateArticleWithRecordID)
		router_api.GET("/article/:record_id", api.GetArticleByRecordID)

		router_api.POST("/record/create", api.CreateRecord)
		router_api.GET("/record/:record_id", api.GetRecordByID)
		router_api.GET("/record/all", api.GetRecordsByUser)
		router_api.POST("/record/:record_id/publish", api.PublishRecord)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.Cfg.ServerPort)
}
