package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuParamLogic {
	return &UpdateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuParamLogic) UpdateMenuParam(in *system.UpdateMenuParamReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, nil
}
