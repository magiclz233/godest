package cache

import (
	"context"
	"encoding/json"
	"time"

	pkgcache "godest/pkg/cache"
)

type JSONStore struct {
	client *pkgcache.RedisClient
}

func NewJSONStore(client *pkgcache.RedisClient) *JSONStore {
	return &JSONStore{client: client}
}

func (s *JSONStore) Get(ctx context.Context, key string, dest any) bool {
	if s == nil || s.client == nil || s.client.Client == nil {
		return false
	}

	val, err := s.client.Client.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	return json.Unmarshal([]byte(val), dest) == nil
}

func (s *JSONStore) Set(ctx context.Context, key string, value any, ttl time.Duration) {
	if s == nil || s.client == nil || s.client.Client == nil {
		return
	}

	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	s.client.Client.Set(ctx, key, data, ttl)
}

func (s *JSONStore) Del(ctx context.Context, keys ...string) {
	if s == nil || s.client == nil || s.client.Client == nil || len(keys) == 0 {
		return
	}

	s.client.Client.Del(ctx, keys...)
}
