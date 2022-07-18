package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID   `json:"uuid" gorm:"comment:user's UUID"`
	Username    string      `json:"userName" gorm:"unique;comment:user's login name'"`
	Password    string      `json:"-"  gorm:"comment:password"`
	NickName    string      `json:"nickName" gorm:"unique;comment:nickname"`
	SideMode    string      `json:"sideMode" gorm:"default:dark;comment:template mode"`
	Avatar      string      `json:"Avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:avatar"`
	BaseColor   string      `json:"baseColor" gorm:"default:#fff;comment:base color of template"`
	ActiveColor string      `json:"activeColor" gorm:"default:#1890ff;comment:active color of template"`
	AuthorityId uint32      `json:"authorityId" gorm:"default:1;comment:user's role id for access control"`
	Authority   Authority   `json:"authority.proto" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:user role id"`
	Authorities []Authority `json:"authorities" gorm:"many2many:user_authorities;"`
	Mobile      string      `json:"mobile"  gorm:"comment:用户手机号"`
	Email       string      `json:"email"  gorm:"comment:用户邮箱"`
	Status      int32       `json:"status" gorm:"default:1;comment:user status 1: normal 2: ban"`
}
