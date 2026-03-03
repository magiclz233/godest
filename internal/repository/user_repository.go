package repository

import (
	"godest/internal/model"
	"godest/pkg/database"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetAll() ([]model.User, error)
}

type GormUserRepository struct{}

func NewUserRepository() UserRepository {
	return &GormUserRepository{}
}

var _ UserRepository = (*GormUserRepository)(nil)

func (r *GormUserRepository) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

func (r *GormUserRepository) GetByUsername(username string) (*model.User, error) {
	var u model.User
	err := database.DB.Where("username = ?", username).First(&u).Error
	return &u, err
}

func (r *GormUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	return users, err
}
