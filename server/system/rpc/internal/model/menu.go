package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuLevel    uint        `json:"-"`
	ParentId     uint32      `json:"parentId" gorm:"comment:parent menu id"`            // parent menu id
	Path         string      `json:"path" gorm:"comment:index path"`                   // index path
	Name         string      `json:"name" gorm:"comment:index name"`                   // index name
	Hidden       bool        `json:"hidden" gorm:"comment:hide the menu"`               // hide menu
	Component    string      `json:"component" gorm:"comment:the path of vue file"`     // the path of vue file
	Sort         int         `json:"sort" gorm:"comment:numbers for sorting"`           // sorting numbers
	MetaData     Meta        `json:"metaData" gorm:"embedded;comment:extra parameters"` // extra parameters
	AuthorityIds []Authority `json:"authorityIds" gorm:"many2many:authority_menus;"`
	Children     []Menu      `json:"children" gorm:"-"`
	Param        []MenuParam `json:"parameters"`
}

type AuthorityMenu struct {
	MenuId      uint            `json:"menuId" gorm:"comment:menu id"`
	AuthorityId uint            `json:"-" gorm:"comment:authority.proto id (role id)"`
	Children    []AuthorityMenu `json:"children" gorm:"-"`
	Parameters  []MenuParam     `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint `json:"btns" gorm:"-"`
}

type Meta struct {
	KeepAlive bool   `json:"keepAlive" gorm:"comment:save in cache?"` // save in cache
	Title     string `json:"title" gorm:"comment:menu name"`          // menu name
	Icon      string `json:"icon" gorm:"comment:menu icon"`           // menu icon
	CloseTab  bool   `json:"closeTab" gorm:"comment:auto close tab"`  // auto close tab
}

type MenuParam struct {
	gorm.Model
	MenuId uint
	Type   string `json:"type" gorm:"comment:pass parameters via params or query "` // pass parameters via params or query
	Key    string `json:"key" gorm:"comment:the key of parameters"`                 // the key of parameters
	Value  string `json:"value" gorm:"comment:the value of parameters"`             // the value of parameters
}
