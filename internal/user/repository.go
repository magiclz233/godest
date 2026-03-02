package user

import "godest/pkg/database"

// Repository 定义用户持久化接口
type Repository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
	GetAll() ([]User, error)
}

// GormRepository 基于 GORM 的用户仓储实现
type GormRepository struct{}

// NewRepository 创建用户仓储
func NewRepository() Repository {
	return &GormRepository{}
}

var _ Repository = (*GormRepository)(nil)

// Create 创建用户
func (r *GormRepository) Create(user *User) error {
	return database.DB.Create(user).Error
}

// GetByUsername 按用户名查询
func (r *GormRepository) GetByUsername(username string) (*User, error) {
	var u User
	err := database.DB.Where("username = ?", username).First(&u).Error
	return &u, err
}

// GetAll 查询全部用户
func (r *GormRepository) GetAll() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error
	return users, err
}
