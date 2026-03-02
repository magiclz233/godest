package router

import (
	"godest/config"
	"godest/internal/platform/http/middleware"
	"godest/internal/user"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Init 初始化 HTTP 路由
func Init(userHandler *user.Handler, jwtUtil *utils.JWTUtil) *gin.Engine {
	gin.SetMode(config.GlobalConfig.App.Mode)

	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/register", userHandler.Register)
		apiV1.POST("/login", userHandler.Login)

		authorized := apiV1.Group("/")
		authorized.Use(middleware.AuthMiddleware(jwtUtil))
		{
			authorized.GET("/users", userHandler.ListUsers)
		}
	}

	return r
}
