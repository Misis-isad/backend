package main

import (
	"context"
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
	err := logging.InitLog()
	if err != nil {
		fmt.Printf("Can't init log: %v", err)
		os.Exit(1)
	}

	logging.Log.Debug("Starting up")

	err = config.LoadConfig()
	if err != nil {
		logging.Log.Fatalf("Can't load config: %v", err)
	}

	db, err := database.InitDb(&config.Cfg)
	if err != nil {
		logging.Log.Fatalf("Can't init db: %v", err)
	}
	if err := db.Pool.Ping(context.Background()); err != nil {
		logging.Log.Fatalf("Can't access db: %v", err)
	}

	apiClient := api.NewApiClient(db)

	gin.SetMode(config.Cfg.GinMode)
	r := gin.Default()

	router_auth := r.Group("/auth")
	{
		router_auth.POST("/user/create", apiClient.CreateUser)
		router_auth.POST("/user/login", apiClient.LoginUser)
	}

	router_api := r.Group("/api/v1")
	router_api.Use(middlewares.JwtAuth())
	{
		router_api.GET("/test", apiClient.TestMiddleware)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.Cfg.ServerPort)
}
