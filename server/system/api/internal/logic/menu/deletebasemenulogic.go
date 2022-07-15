package menu

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBaseMenuLogic {
	return &DeleteBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBaseMenuLogic) DeleteBaseMenu(req *types.IdReq) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
