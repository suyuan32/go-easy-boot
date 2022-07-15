package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityByIdLogic {
	return &GetMenuAuthorityByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuAuthorityByIdLogic) GetMenuAuthorityById(in *system.IDReq) (*system.MenuAuthorityResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuAuthorityResp{}, nil
}
