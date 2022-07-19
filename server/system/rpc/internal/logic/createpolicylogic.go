package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePolicyLogic {
	return &CreatePolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePolicyLogic) CreatePolicy(in *system.CreatePolicyReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
