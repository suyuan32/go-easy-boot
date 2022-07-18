package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePolicyLogic {
	return &UpdatePolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  casbin service
func (l *UpdatePolicyLogic) UpdatePolicy(in *system.UpdatePolicyReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
