package router

import (
	"godest/internal/handler"
	"godest/internal/transport/http/middleware"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(apiV1 *gin.RouterGroup, userHandler *handler.UserHandler, jwtUtil *utils.JWTUtil) {
	apiV1.POST("/register", userHandler.Register)
	apiV1.POST("/login", userHandler.Login)

	authorized := apiV1.Group("/")
	authorized.Use(middleware.AuthMiddleware(jwtUtil))
	authorized.GET("/users", userHandler.ListUsers)
}
