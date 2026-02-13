package model

import "gorm.io/gorm"

// User 用户模型
// User model structure
type User struct {
	gorm.Model        // 内嵌 gorm.Model，包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"uniqueIndex;size:100;not null" json:"username"` // 用户名，唯一索引，非空
	Email      string `gorm:"uniqueIndex;size:100;not null" json:"email"`    // 邮箱，唯一索引，非空
	Password   string `gorm:"not null" json:"-"`                             // 密码，非空，JSON 输出时忽略
}

// TableName 指定表名
// TableName overrides the table name
func (User) TableName() string {
	return "users"
}
