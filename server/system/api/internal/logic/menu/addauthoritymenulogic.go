package menu

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAuthorityMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAuthorityMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAuthorityMenuLogic {
	return &AddAuthorityMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAuthorityMenuLogic) AddAuthorityMenu(req *types.AuthorityMenu) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
