// Code generated by goctl. DO NOT EDIT.
package types

type AuthorityMenu struct {
	BaseMenu
	MenuId      string          `json:"menuId"`
	AuthorityId uint            `json:"-"`
	Children    []AuthorityMenu `json:"children"`
	Param       []MenuParam     `json:"param"`
	Btns        map[string]uint `json:"btns"`
}

type BaseMenu struct {
	BaseInfo
	MenuLevel    uint        `json:"-"`
	ParentId     string      `json:"parentId"`  // parent menu id
	Path         string      `json:"path"`      // index path
	Name         string      `json:"name"`      // index name
	Hidden       bool        `json:"hidden"`    // hide menu
	Component    string      `json:"component"` // the path of vue file
	Sort         int         `json:"sort"`      // sorting numbers
	MetaData     Meta        `json:"metaData"`  // extra parameters
	AuthorityIds []Authority `json:"authorityIds"`
	Children     []BaseMenu  `json:"children"`
	Param        []MenuParam `json:"parameters"`
}

type Meta struct {
	KeepAlive bool   `json:"keepAlive"` // save in cache
	Title     string `json:"title"`     // menu name
	Icon      string `json:"icon"`      // menu icon
	CloseTab  bool   `json:"closeTab"`  // auto close tab
}

type MenuParam struct {
	BaseInfo
	MenuId uint   `json:"menuId"`
	Type   string `json:"type"`  // pass parameters via params or query
	Key    string `json:"key"`   // the key of parameters
	Value  string `json:"value"` // the value of parameters
}

type Authority struct {
	BaseInfo
	AuthorityId   uint32          `json:"authorityId"`
	AuthorityName string          `json:"authorityName"`
	ParentId      uint32          `json:"parentId"`
	AuthorityIds  []*Authority    `json:"AuthorityIds"`
	Children      []Authority     `json:"children"`
	Menu          []AuthorityMenu `json:"menus" gorm:"many2many:authority_menus;"`
	DefaultRouter string          `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
}

type AuthorityMenuResp struct {
	BaseMsg
	Data AuthorityMenu `json:"data"`
}

type MenuListResp struct {
	PageList
	Data []AuthorityMenu `json:"data"`
}

type BaseMenuListResp struct {
	PageList
	Data []BaseMenu `json:"data"`
}

type BaseMenuResp struct {
	BaseInfo
	Data BaseMenu `json:"data"`
}

type BaseMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type PageInfo struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Keyword  string `form:"keyword"`
}

type PageList struct {
	Code     int32    `json:"code"`
	Total    int64    `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
	Data     []string `json:"list"`
}

type BaseInfo struct {
	ID        uint   `json:"ID"`
	CreatedAt uint64 `json:"createdAt"`
	UpdatedAt uint64 `json:"updatedAt"`
	DeletedAt uint64 `json:"deletedAt"`
}

type IdReq struct {
	ID uint `json:"ID"`
}

type LoginReq struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	CaptchaId int32  `form:"captchaId"`
	Captcha   string `form:"captcha"`
}

type LoginResp struct {
	BaseMsg
	Data LoginRespData `json:"data"`
}

type LoginRespData struct {
	UserId       int64  `json:"userId"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	RoleId       int32  `json:"roleId"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterReq struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	CaptchaId string `form:"captchaId"`
	Captcha   string `form:"captcha"`
	Email     string `form:"email"`
}

type RegisterResp struct {
	BaseMsg
	Data RegisterRespData `json:"data"`
}

type RegisterRespData struct {
	Status string `json:"status"`
}

type ChangePasswordReq struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	NewPassword string `form:"newPassword"`
}

type ChangePasswordResp struct {
	BaseMsg
}

type ModifyInfoReq struct {
	UserId       int64  `form:"userId"`
	Nickname     string `form:"nickname"`
	Mobile       string `form:"mobile"`
	AuthorityIds string `form:"authorityIds"`
	Email        string `form:"email"`
	Avatar       string `form:"avatar"`
	SideMode     string `form:"sideMode"`
	Status       int32  `form:"status"`
}

type UserInfoResp struct {
	UserId       int64  `json:"userId"`
	UUID         int64  `json:"uuid"`
	Username     int64  `json:"username"`
	Nickname     string `json:"nickname"`
	Mobile       string `json:"mobile"`
	AuthorityIds string `json:"authorityIds"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	SideMode     string `json:"sideMode"`
	Status       int32  `json:"status"`
	CreateAt     uint64 `json:"createAt"`
	UpdateAt     uint64 `json:"updateAt"`
	DeleteAt     uint64 `json:"deleteAt"`
}

type UserInfoListResp struct {
	PageList
	Data []UserInfoResp `json:"data"`
}
