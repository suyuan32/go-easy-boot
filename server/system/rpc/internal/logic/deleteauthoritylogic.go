package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAuthorityLogic {
	return &DeleteAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAuthorityLogic) DeleteAuthority(in *system.IDReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
