package authority

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyAuthorityLogic {
	return &CopyAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyAuthorityLogic) CopyAuthority(req *types.IdReq) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
