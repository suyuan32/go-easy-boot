package menu

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBaseMenuLogic {
	return &CreateBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBaseMenuLogic) CreateBaseMenu(req *types.BaseMenu) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
