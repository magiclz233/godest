package repository

import (
	"godest/internal/model"
	"godest/pkg/database"
)

type TenantRepository interface {
	Create(user *model.Tenant) error
	GetByDomain(domain string) (*model.Tenant, error)
	GetAll() ([]model.Tenant, error)
	GetByID(id uint) (*model.Tenant, error)
}

type GormTenantRepository struct{}

var _ TenantRepository = (*GormTenantRepository)(nil)

// Create implements [TenantRepository].
func (g *GormTenantRepository) Create(user *model.Tenant) error {
	return database.DB.Create(user).Error
}

// GetAll implements [TenantRepository].
func (g *GormTenantRepository) GetAll() ([]model.Tenant, error) {
	var tenants []model.Tenant
	if err := database.DB.Where("del_flag != ?", 1).Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

// GetByID implements [TenantRepository].
func (g *GormTenantRepository) GetByID(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := database.DB.Where("id = ? AND del_flag != ?", id, 1).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

// GetByDomain implements [TenantRepository].
func (g *GormTenantRepository) GetByDomain(domain string) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := database.DB.Where("domain = ?", domain).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func NewTenantRepository() TenantRepository {
	return &GormTenantRepository{}
}


