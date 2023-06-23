package main

import (
	"fmt"
	"os"
	"profbuh/internal/config"
	"profbuh/internal/database"
	"profbuh/internal/logging"
	"profbuh/internal/router"

	_ "profbuh/docs"

	"github.com/gin-gonic/gin"
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
	logging.Log.Info("Db connected")

	gin.SetMode(config.Cfg.GinMode)
	r := router.NewRouter(db)
	logging.Log.Info("Set up router")

	r.Run(":" + config.Cfg.ServerPort)
}
