package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	NameEn      string `gorm:"type:varchar(100);not null" json:"name_en"`
	Code        string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	Description string `gorm:"type:text" json:"description"`
	Status      int    `gorm:"type:tinyint;default:1;comment:状态 0-禁用 1-启用"`
	DelFlag     int    `gorm:"type:tinyint;default:0;comment:删除标志 0-未删除 1-已删除"`
	TenantId    uint   `gorm:"not null;comment:租户ID" json:"tenant_id"`
	CreatorId   uint   `gorm:"not null;comment:创建者ID" json:"creator_id"`
	Sort        int    `gorm:"type:int;default:0;comment:排序" json:"sort"`
	IsSystem    bool   `gorm:"default:false" json:"is_system"` // 是否系统内置角色
}
