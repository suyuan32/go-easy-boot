package menu

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseMenuLogic {
	return &UpdateBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBaseMenuLogic) UpdateBaseMenu(req *types.BaseMenu) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
