package main

import (
	"godest/config"
	"godest/internal/model"
	"godest/pkg/database"
	"godest/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置
	config.LoadConfig()

	// 2. 初始化日志
	logger.Init()
	logger.Log.Info("Starting GoPro Application...")

	// 3. 初始化数据库
	database.Init()

	// 4. 自动迁移数据库表
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		logger.Log.Fatal("Database migration failed", zap.Error(err))
	}
	logger.Log.Info("Database connected and migrated")

	// 5. 使用 Wire 初始化应用 (依赖注入)
	r, err := InitApp()
	if err != nil {
		logger.Log.Fatal("Failed to initialize app", zap.Error(err))
	}

	// 6. 启动服务
	port := config.GlobalConfig.App.Port
	logger.Log.Info("Server starting", zap.String("port", port))
	if err := r.Run(port); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
