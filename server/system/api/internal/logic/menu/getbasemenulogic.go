package menu

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseMenuLogic {
	return &GetBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseMenuLogic) GetBaseMenu(req *types.IdReq) (resp *types.BaseMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
