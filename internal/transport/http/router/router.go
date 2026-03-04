package router

import (
	"godest/internal/config"
	"godest/internal/handler"
	"godest/internal/transport/http/middleware"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(handlers *handler.Handlers, jwtUtil *utils.JWTUtil) *gin.Engine {
	gin.SetMode(config.GlobalConfig.App.Mode)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	apiV1 := r.Group("/api/v1")
	registerUserRoutes(apiV1, handlers.User, jwtUtil)
	return r
}
