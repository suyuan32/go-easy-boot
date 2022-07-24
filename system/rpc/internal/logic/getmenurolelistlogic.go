package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRoleListLogic {
	return &GetMenuRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuRoleListLogic) GetMenuRoleList(in *system.PageInfoReq) (*system.MenuRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuRoleListResp{}, nil
}
