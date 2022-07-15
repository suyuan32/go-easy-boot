package menu

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseMenuTreeLogic {
	return &GetBaseMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseMenuTreeLogic) GetBaseMenuTree() (resp *types.BaseMenuListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
