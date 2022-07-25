package user

import (
	"context"
	"errors"
	"system/api/internal/svc"
	"system/api/internal/types"
	"system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	if l.ctx.Value("userId").(string) == "" {
		return nil, errors.New("not log in yet")
	}
	user, err := l.svcCtx.SystemRpc.GetUserById(context.Background(),
		&system.UUIDReq{UUID: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		UserId:   user.UUID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Roles: types.GetUserRoleInfo{
			RoleName: user.RoleName,
			Value:    user.RoleValue,
		},
	}, nil
}
