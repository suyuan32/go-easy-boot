package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPolicyByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyByRoleIdLogic {
	return &GetPolicyByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPolicyByRoleIdLogic) GetPolicyByRoleId(in *system.IDReq) (*system.UpdatePolicyReq, error) {
	// todo: add your logic here and delete this line

	return &system.UpdatePolicyReq{}, nil
}
