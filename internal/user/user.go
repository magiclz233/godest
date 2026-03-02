package user

import "gorm.io/gorm"

// User 用户实体
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
