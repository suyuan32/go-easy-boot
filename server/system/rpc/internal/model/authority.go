package model

import (
	"gorm.io/gorm"
)

type Authority struct {
	gorm.Model
	AuthorityId   uint64          `json:"authorityId" gorm:"not null;unique;primary_key;comment:role id for access control;size:90"`
	AuthorityName string          `json:"authorityName" gorm:"comment:role name"`
	ParentId      uint64          `json:"parentId" gorm:"comment:parent id"`
	AuthorityIds  []*Authority    `json:"AuthorityIds" gorm:"many2many:authority_ids;"`
	Children      []Authority     `json:"children" gorm:"-"`
	Menu          []AuthorityMenu `json:"menus" gorm:"many2many:authority_menus;"`
	DefaultRouter string          `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
}
