package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	NameEn     string `gorm:"type:varchar(100);not null" json:"name_en"`
	Permission string `gorm:"type:varchar(100);not null" json:"permission"`
	Path       string `gorm:"type:varchar(255);not null" json:"path"`
	ParentId   uint   `gorm:"default:0;comment:父级ID" json:"parent_id"`
	Icon       string `gorm:"type:varchar(255);comment:菜单图标" json:"icon,omitempty"`
	Status     int    `gorm:"type:tinyint;default:1;comment:状态 0-禁用 1-启用" json:"status"`
	DelFlag    int    `gorm:"type:tinyint;default:0;comment:删除标志 0-未删除 1-已删除" json:"del_flag"`
	Sort       int    `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Type       int    `gorm:"type:tinyint;default:0;comment:菜单类型 0-菜单 1-按钮" json:"type"`
	Grouping   string `gorm:"type:varchar(100);comment:权限分组" json:"grouping,omitempty"`
	Remark     string `gorm:"type:varchar(255);comment:备注" json:"remark,omitempty"`
	TenantId   uint   `gorm:"not null;comment:租户ID" json:"tenant_id"`
}
