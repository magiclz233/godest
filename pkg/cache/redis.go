package cache

import (
	"context"
	"fmt"
	"log"

	"godest/internal/config"

	"github.com/redis/go-redis/v9"
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
		log.Printf("warning: failed to connect to redis: %v", err)
	} else {
		fmt.Println("redis connected")
	}

	return &RedisClient{Client: rdb}
}
