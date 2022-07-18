package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuAuthorityLogic {
	return &CreateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// menu authority management
func (l *CreateMenuAuthorityLogic) CreateMenuAuthority(in *system.CreateMenuAuthorityReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
