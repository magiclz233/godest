package cache

import (
	"context"
	"fmt"
	"godest/config"
	"log"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis 客户端封装
// RedisClient wrapper
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient 创建 Redis 客户端
// NewRedisClient creates a new Redis client
func NewRedisClient() *RedisClient {
	cfg := config.GlobalConfig.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		// 在生产环境中，Redis 连接失败可能不应该直接 panic，而是降级或重试
		// 这里为了演示方便直接打印错误
		log.Printf("Warning: Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Redis connected successfully")
	}

	return &RedisClient{Client: rdb}
}
