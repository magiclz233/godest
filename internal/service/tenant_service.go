package service

import (
	"context"

	internalcache "godest/internal/cache"
	"godest/internal/model"
	"godest/internal/repository"
	pkgcache "godest/pkg/cache"
)

type TenantService struct {
	repo  repository.TenantRepository
	cache *internalcache.TenantCache
}

func NewTenantService(
	repo repository.TenantRepository,
	redis *pkgcache.RedisClient,
) *TenantService {
	return &TenantService{
		repo:  repo,
		cache: internalcache.NewTenantCache(redis),
	}
}

func (s *TenantService) GetAll() ([]model.Tenant, error) {
	ctx := context.Background()

	if tenants, ok := s.cache.GetAll(ctx); ok {
		return tenants, nil
	}

	tenants, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	s.cache.SetAll(ctx, tenants)
	return tenants, nil
}

func (s *TenantService) GetByID(id uint) (*model.Tenant, error) {
	ctx := context.Background()

	if tenant, ok := s.cache.GetByID(ctx, id); ok {
		return tenant, nil
	}

	tenant, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	s.cache.Set(ctx, tenant)
	return tenant, nil
}

func (s *TenantService) GetByDomain(domain string) (*model.Tenant, error) {
	ctx := context.Background()

	if tenant, ok := s.cache.GetByDomain(ctx, domain); ok {
		return tenant, nil
	}

	tenant, err := s.repo.GetByDomain(domain)
	if err != nil {
		return nil, err
	}

	s.cache.Set(ctx, tenant)
	return tenant, nil
}

func (s *TenantService) Create(tenant *model.Tenant) error {
	if err := s.repo.Create(tenant); err != nil {
		return err
	}

	ctx := context.Background()
	s.cache.Set(ctx, tenant)
	s.cache.InvalidateList(ctx)
	return nil
}
