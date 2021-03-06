package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleId        uint64 `json:"roleId" gorm:"unique;comment:role id"`
	Name          string `json:"roleName" gorm:"comment:role name"`
	Value         string `json:"value" gorm:"comment: role value for permission control in front end"`
	ParentId      uint64 `json:"parentId" gorm:"comment:parent id"`
	Children      []Role `json:"children" gorm:"-"`
	Menu          []Menu `json:"menus" gorm:"many2many:role_menus;"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
}
