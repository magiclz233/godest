package repository

import (
	"go_pro/internal/model"
	"go_pro/pkg/database"
)

// IUserRepository 用户仓库接口
// IUserRepository defines the interface for user repository
type IUserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetAll() ([]model.User, error)
}

// UserRepository 用户数据访问层
// UserRepository handles database operations for users
type UserRepository struct{}

// NewUserRepository 创建 UserRepository 实例
// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

// Ensure UserRepository implements IUserRepository
var _ IUserRepository = (*UserRepository)(nil)

// Create 创建用户
// Create inserts a new user into the database
func (r *UserRepository) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

// GetByUsername 根据用户名查找用户
// GetByUsername finds a user by username
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	// 使用 First 查询单条记录
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetAll 获取所有用户
// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	return users, err
}
