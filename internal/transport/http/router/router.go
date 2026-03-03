package router

import (
	"godest/internal/config"
	"godest/internal/handler"
	"godest/internal/transport/http/middleware"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *handler.UserHandler, jwtUtil *utils.JWTUtil) *gin.Engine {
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
