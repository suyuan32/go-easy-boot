package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityListLogic {
	return &GetMenuAuthorityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuAuthorityListLogic) GetMenuAuthorityList(in *system.PageInfoReq) (*system.MenuAuthorityListResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuAuthorityListResp{}, nil
}
