package model

import (
	"gorm.io/gorm"
)

type Authority struct {
	gorm.Model
	Name          string      `json:"authorityName" gorm:"comment:role name"`
	ParentId      uint64      `json:"parentId" gorm:"comment:parent id"`
	Children      []Authority `json:"children" gorm:"-"`
	Menu          []Menu      `json:"menus" gorm:"many2many:authority_menus;"`
	DefaultRouter string      `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
}
