package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Birthday time.Time `json:"birthday"`
	Phone    string `gorm:"size:20" json:"phone,omitempty"`
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
