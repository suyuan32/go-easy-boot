package authority

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAuthorityLogic {
	return &CreateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAuthorityLogic) CreateAuthority(req *types.Authority) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
