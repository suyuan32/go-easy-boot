package logic

import (
	"context"
	"google.golang.org/grpc/codes"

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
	policy, err := l.svcCtx.Role.AddPolicy(in.AuthorityId, in.Info.Path, in.Info.Method)
	if err != nil {
		return nil, err
	}
	if policy == true {
		return &system.BaseResp{
			Code: uint32(codes.OK),
			Msg:  "ok",
		}, nil
	} else {
		return &system.BaseResp{
			Code: uint32(codes.AlreadyExists),
			Msg:  "exists",
		}, nil
	}
}
