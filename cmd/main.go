package main

import (
	"godest/internal/app"
	"godest/internal/config"
	"godest/internal/model"
	"godest/pkg/database"
	"godest/pkg/log"

	"go.uber.org/zap"
)

func main() {
	log.Init()
	config.LoadConfig()
	log.Info("starting godest")

	database.Init()
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("database migration failed", zap.Error(err))
	}

	// Initialize the application and its dependencies
	application := app.NewApp()

	port := config.GlobalConfig.App.Port
	log.Info("server starting", zap.String("port", port))
	if err := application.Run(port); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
