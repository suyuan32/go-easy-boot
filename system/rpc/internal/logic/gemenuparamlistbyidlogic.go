package logic

import (
	"context"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeMenuParamListByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeMenuParamListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeMenuParamListByIdLogic {
	return &GeMenuParamListByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeMenuParamListByIdLogic) GeMenuParamListById(in *system.IDReq) (*system.MenuParamListResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuParamListResp{}, nil
}
