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
	NickName    string      `json:"nickName" gorm:"unique;default:系统用户;comment:nickname"`
	SideMode    string      `json:"sideMode" gorm:"default:dark;comment:template mode"` // user's template mode
	Avatar      string      `json:"Avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:avatar"`
	BaseColor   string      `json:"baseColor" gorm:"default:#fff;comment:base color of template"`
	ActiveColor string      `json:"activeColor" gorm:"default:#1890ff;comment:active color of template"`    // 活跃颜色
	AuthorityId uint        `json:"authorityId" gorm:"default:1;comment:user's role id for access control"` // 用户角色ID
	Authority   Authority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []Authority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Mobile      string      `json:"mobile"  gorm:"comment:用户手机号"`                    // 用户手机号
	Email       string      `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Status      int         `json:"status" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}
