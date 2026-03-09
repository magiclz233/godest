package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Nickname string    `gorm:"not null;comment:用户昵称" json:"nickname"`
	TenantId uint      `gorm:"not null;comment:租户ID" json:"tenant_id"`
	Birthday time.Time `json:"birthday"`
	DeptID   uint      `json:"dept_id,omitempty"`
	Phone    string    `gorm:"size:20" json:"phone,omitempty"`
	Email    string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Gender   int       `gorm:"type:tinyint;default:0;comment:性别 0-未知 1-男 2-女"`
	Avatar   string    `gorm:"type:varchar(255);comment:用户头像"`
	Password string    `gorm:"not null" json:"-"`
	Status   int       `gorm:"type:tinyint;default:1;comment:用户状态 0-禁用 1-启用"`
	DelFlag  int       `gorm:"type:tinyint;default:0;comment:删除标志 0-未删除 1-已删除"`
	Metadata string    `gorm:"type:jsonb" json:"metadata,omitempty"` // 存储用户扩展属性
}

func (User) TableName() string {
	return "users"
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
