package main

import (
	"godest/internal/platform/http/router"
	"godest/internal/user"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

// InitApp 手写依赖注入并组装应用
func InitApp() (*gin.Engine, error) {
	repo := user.NewRepository()
	redisClient := cache.NewRedisClient()
	jwtUtil := utils.NewJWTUtil()
	passwordUtil := utils.NewPasswordUtil()

	service := user.NewService(repo, redisClient, jwtUtil, passwordUtil)
	handler := user.NewHandler(service)
	engine := router.Init(handler, jwtUtil)
	return engine, nil
}
