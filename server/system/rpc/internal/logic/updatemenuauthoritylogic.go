package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuAuthorityLogic {
	return &UpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuAuthorityLogic) UpdateMenuAuthority(in *system.UpdateMenuAuthorityReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
