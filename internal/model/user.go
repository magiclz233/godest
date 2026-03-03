package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

func (User) TableName() string {
	return "users"
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
