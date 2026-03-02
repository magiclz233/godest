package repository

import (
	"godest/internal/user/domain"
	"godest/pkg/database"
)

// UserRepository 用户仓储实现
type UserRepository struct{}

// NewUserRepository 创建用户仓储
func NewUserRepository() domain.Repository {
	return &UserRepository{}
}

var _ domain.Repository = (*UserRepository)(nil)

// Create 创建用户
func (r *UserRepository) Create(user *domain.User) error {
	return database.DB.Create(user).Error
}

// GetByUsername 按用户名查询
func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetAll 查询全部用户
func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := database.DB.Find(&users).Error
	return users, err
}
