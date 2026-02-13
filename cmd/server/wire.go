//go:build wireinject
// +build wireinject

package main

import (
	"godest/internal/handler"
	"godest/internal/repository"
	"godest/internal/service"
	"godest/pkg/cache"
	"godest/pkg/utils"
	"godest/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() (*gin.Engine, error) {
	wire.Build(
		repository.NewUserRepository,
		cache.NewRedisClient,
		utils.NewJWTUtil,
		utils.NewPasswordUtil,
		service.NewUserService,
		handler.NewUserHandler,
		router.Init,
	)
	return &gin.Engine{}, nil
}
