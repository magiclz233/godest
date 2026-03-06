package cache

import (
	"context"

	"godest/internal/config"
	"godest/pkg/log"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	cfg := config.GlobalConfig.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Warn("failed to connect to redis", zap.Error(err))
	} else {
		log.Info("redis connected")
	}

	return &RedisClient{Client: rdb}
}
