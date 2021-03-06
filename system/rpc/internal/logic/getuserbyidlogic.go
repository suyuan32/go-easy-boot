package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"

	"system/rpc/internal/model"
	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *system.UUIDReq) (*system.UserInfoResp, error) {
	var u model.User
	result := l.svcCtx.DB.Where("uuid = ?", in.UUID).First(&u)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	roleName, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId))
	roleValue, err := l.svcCtx.Redis.Hget("roleData", roleName)
	if err != nil {
		return nil, err
	}
	return &system.UserInfoResp{
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		RoleId:    u.RoleId,
		RoleName:  roleName,
		RoleValue: roleValue,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Status:    u.Status,
		UserId:    uint64(u.ID),
		Username:  u.Username,
		UUID:      u.UUID,
		CreateAt:  u.CreatedAt.Unix(),
		UpdateAt:  u.UpdatedAt.Unix(),
		DeleteAt:  u.DeletedAt.Time.Unix(),
	}, nil
}
