//go:build wireinject
// +build wireinject

package main

import (
	httpRouter "godest/internal/platform/http/router"
	"godest/internal/user/api"
	"godest/internal/user/app"
	"godest/internal/user/infra/repository"
	"godest/pkg/cache"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() (*gin.Engine, error) {
	wire.Build(
		repository.NewUserRepository,
		cache.NewRedisClient,
		utils.NewJWTUtil,
		utils.NewPasswordUtil,
		app.NewUserService,
		api.NewUserHandler,
		httpRouter.Init,
	)
	return &gin.Engine{}, nil
}
