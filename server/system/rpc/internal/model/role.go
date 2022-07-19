package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleId        uint64 `json:"roleId" gorm:"unique;comment:role id (role id)"`
	Name          string `json:"roleName" gorm:"comment:role name"`
	ParentId      uint64 `json:"parentId" gorm:"comment:parent id"`
	Children      []Role `json:"children" gorm:"-"`
	Menu          []Menu `json:"menus" gorm:"many2many:role_menus;"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
}
