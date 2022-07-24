package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRoleByIdLogic {
	return &GetMenuRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuRoleByIdLogic) GetMenuRoleById(in *system.IDReq) (*system.MenuRoleResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuRoleResp{}, nil
}
