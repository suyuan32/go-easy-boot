package menu

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleMenuLogic {
	return &AddRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleMenuLogic) AddRoleMenu(req *types.RoleMenu) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
