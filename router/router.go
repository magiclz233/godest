package router

import (
	"go_pro/config"
	"go_pro/internal/handler"
	"go_pro/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
// Init initializes the router
func Init(userHandler *handler.UserHandler) *gin.Engine {
	// 设置 Gin 模式
	gin.SetMode(config.GlobalConfig.App.Mode)

	r := gin.Default()

	// 定义路由组
	api := r.Group("/api/v1")
	{
		// 公开路由
		api.POST("/register", userHandler.Register) // 注册
		api.POST("/login", userHandler.Login)       // 登录

		// 需要认证的路由
		authorized := api.Group("/")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.GET("/users", userHandler.ListUsers) // 获取用户列表
		}
	}

	return r
}
