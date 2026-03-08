package cache

import (
	"context"
	"fmt"
	"time"

	"godest/internal/model"
	pkgcache "godest/pkg/cache"
)

const (
	tenantListCacheKey    = "tenants:all"
	tenantDetailCacheTmpl = "tenant:%d"
	tenantDomainCacheTmpl = "tenant:domain:%s"
	tenantCacheTTL        = time.Minute
)

type TenantCache struct {
	store *JSONStore
}

func NewTenantCache(client *pkgcache.RedisClient) *TenantCache {
	return &TenantCache{store: NewJSONStore(client)}
}

func (c *TenantCache) GetAll(ctx context.Context) ([]model.Tenant, bool) {
	var tenants []model.Tenant
	if !c.store.Get(ctx, tenantListCacheKey, &tenants) {
		return nil, false
	}
	return tenants, true
}

func (c *TenantCache) SetAll(ctx context.Context, tenants []model.Tenant) {
	c.store.Set(ctx, tenantListCacheKey, tenants, tenantCacheTTL)
}

func (c *TenantCache) GetByID(ctx context.Context, id uint) (*model.Tenant, bool) {
	return c.getOne(ctx, fmt.Sprintf(tenantDetailCacheTmpl, id))
}

func (c *TenantCache) GetByDomain(ctx context.Context, domain string) (*model.Tenant, bool) {
	return c.getOne(ctx, fmt.Sprintf(tenantDomainCacheTmpl, domain))
}

func (c *TenantCache) Set(ctx context.Context, tenant *model.Tenant) {
	if tenant == nil {
		return
	}

	c.store.Set(ctx, fmt.Sprintf(tenantDetailCacheTmpl, tenant.ID), tenant, tenantCacheTTL)
	c.store.Set(ctx, fmt.Sprintf(tenantDomainCacheTmpl, tenant.Domain), tenant, tenantCacheTTL)
}

func (c *TenantCache) InvalidateList(ctx context.Context) {
	c.store.Del(ctx, tenantListCacheKey)
}

func (c *TenantCache) Invalidate(ctx context.Context, tenant *model.Tenant) {
	if tenant == nil {
		return
	}

	c.store.Del(
		ctx,
		fmt.Sprintf(tenantDetailCacheTmpl, tenant.ID),
		fmt.Sprintf(tenantDomainCacheTmpl, tenant.Domain),
	)
}

func (c *TenantCache) getOne(ctx context.Context, key string) (*model.Tenant, bool) {
	var tenant model.Tenant
	if !c.store.Get(ctx, key, &tenant) {
		return nil, false
	}
	return &tenant, true
}
