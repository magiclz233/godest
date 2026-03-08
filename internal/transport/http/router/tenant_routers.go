package router

import (
	"godest/internal/handler"
	"godest/internal/transport/http/middleware"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

func registerTenantRoutes(apiV1 *gin.RouterGroup,
	tenantHandler *handler.TenantHandler,
	jwtUtil *utils.JWTUtil) {
	authorized := apiV1.Group("/tenants")
	authorized.Use(middleware.AuthMiddleware(jwtUtil))
	authorized.GET("/all", tenantHandler.GetAll)
	authorized.GET("/:id", tenantHandler.GetByID)
}
