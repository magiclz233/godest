package model

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name 	  string `gorm:"type:varchar(100);not null" json:"name"`
	NameEn	string `gorm:"type:varchar(100);not null" json:"name_en"`
	Domain	string `gorm:"type:varchar(255);not null" json:"domain"`
	Status  int    `gorm:"type:tinyint;default:1;comment:状态 0-禁用 1-启用" json:"status"`
	DelFlag int    `gorm:"type:tinyint;default:0;comment:删除标志 0-未删除 1-已删除" json:"del_flag"`
	Lang 	string `gorm:"type:varchar(20);default:'zh-CN';comment:默认语言" json:"lang"`
	InitPwd  string `gorm:"type:varchar(255);not null;comment:初始密码" json:"init_pwd"`
}