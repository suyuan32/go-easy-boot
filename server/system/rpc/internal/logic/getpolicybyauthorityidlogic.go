package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyByAuthorityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPolicyByAuthorityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyByAuthorityIdLogic {
	return &GetPolicyByAuthorityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPolicyByAuthorityIdLogic) GetPolicyByAuthorityId(in *system.IDReq) (*system.UpdatePolicyReq, error) {
	// todo: add your logic here and delete this line

	return &system.UpdatePolicyReq{}, nil
}
