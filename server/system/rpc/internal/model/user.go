package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" gorm:"comment:user's UUID"`
	Username    string    `json:"username" gorm:"unique;comment:user's login name'"`
	Password    string    `json:"-"  gorm:"comment:password"`
	Nickname    string    `json:"nickName" gorm:"unique;comment:nickname"`
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:template mode"`
	Avatar      string    `json:"Avatar" gorm:"comment:avatar"`
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:base color of template"`
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:active color of template"`
	AuthorityId uint32    `json:"authorityId" gorm:"default:2;comment:user's role id for access control"`
	Authority   Authority `json:"authority" gorm:"foreignKey: AuthorityId;comment:user role id"`
	Mobile      string    `json:"mobile"  gorm:"comment:mobile number"`
	Email       string    `json:"email"  gorm:"comment:email address"`
	Status      int32     `json:"status" gorm:"default:1;comment:user status 1: normal 2: ban"`
}
