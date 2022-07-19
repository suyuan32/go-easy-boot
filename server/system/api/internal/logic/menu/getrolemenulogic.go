package menu

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenuLogic {
	return &GetRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleMenuLogic) GetRoleMenu(req *types.IdReq) (resp *types.RoleMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
