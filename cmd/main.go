package main

import (
	"godest/internal/app"
	"godest/internal/config"
	"godest/internal/model"
	"godest/pkg/database"
	"godest/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	config.LoadConfig()

	logger.Init()
	logger.Log.Info("starting godest")

	database.Init()
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		logger.Log.Fatal("database migration failed", zap.Error(err))
	}

	// Initialize the application and its dependencies
	application := app.NewApp()

	port := config.GlobalConfig.App.Port
	logger.Log.Info("server starting", zap.String("port", port))
	if err := application.Run(port); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}
}
