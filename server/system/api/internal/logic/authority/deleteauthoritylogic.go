package authority

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAuthorityLogic {
	return &DeleteAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAuthorityLogic) DeleteAuthority(req *types.IdReq) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
