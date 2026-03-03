package main

import (
	"godest/internal/config"
	"godest/internal/handler"
	"godest/internal/model"
	"godest/internal/repository"
	"godest/internal/service"
	"godest/internal/transport/http/router"
	"godest/pkg/cache"
	"godest/pkg/database"
	"godest/pkg/logger"
	"godest/pkg/utils"

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

	userRepo := repository.NewUserRepository()
	redisClient := cache.NewRedisClient()
	jwtUtil := utils.NewJWTUtil()
	passwordUtil := utils.NewPasswordUtil()

	userService := service.NewUserService(userRepo, redisClient, jwtUtil, passwordUtil)
	userHandler := handler.NewUserHandler(userService)
	engine := router.NewRouter(userHandler, jwtUtil)

	port := config.GlobalConfig.App.Port
	logger.Log.Info("server starting", zap.String("port", port))
	if err := engine.Run(port); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}
}
