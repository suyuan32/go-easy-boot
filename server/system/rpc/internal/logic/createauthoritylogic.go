package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAuthorityLogic {
	return &CreateAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  authority service
func (l *CreateAuthorityLogic) CreateAuthority(in *system.CreateAuthorityReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
