package cache

import (
	"context"
	"fmt"
	"time"

	"godest/internal/model"
	pkgcache "godest/pkg/cache"
)

const (
	userListCacheKey    = "users:all"
	userDetailCacheTmpl = "user:%d"
	userCacheTTL        = time.Minute
)

type UserCache struct {
	store *JSONStore
}

func NewUserCache(client *pkgcache.RedisClient) *UserCache {
	return &UserCache{store: NewJSONStore(client)}
}

func (c *UserCache) GetAll(ctx context.Context) ([]model.User, bool) {
	var users []model.User
	if !c.store.Get(ctx, userListCacheKey, &users) {
		return nil, false
	}
	return users, true
}

func (c *UserCache) SetAll(ctx context.Context, users []model.User) {
	c.store.Set(ctx, userListCacheKey, users, userCacheTTL)
}

func (c *UserCache) GetByID(ctx context.Context, id uint) (*model.User, bool) {
	var user model.User
	if !c.store.Get(ctx, fmt.Sprintf(userDetailCacheTmpl, id), &user) {
		return nil, false
	}
	return &user, true
}

func (c *UserCache) Set(ctx context.Context, user *model.User) {
	if user == nil {
		return
	}

	c.store.Set(ctx, fmt.Sprintf(userDetailCacheTmpl, user.ID), user, userCacheTTL)
}

func (c *UserCache) InvalidateList(ctx context.Context) {
	c.store.Del(ctx, userListCacheKey)
}

func (c *UserCache) Invalidate(ctx context.Context, user *model.User) {
	if user == nil {
		return
	}

	c.store.Del(ctx, fmt.Sprintf(userDetailCacheTmpl, user.ID))
}
