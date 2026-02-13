//go:build wireinject
// +build wireinject

package main

import (
	"go_pro/internal/handler"
	"go_pro/internal/repository"
	"go_pro/internal/service"
	"go_pro/pkg/cache"
	"go_pro/pkg/utils"
	"go_pro/router"

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
