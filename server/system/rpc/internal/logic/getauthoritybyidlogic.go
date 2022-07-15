package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorityByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityByIdLogic {
	return &GetAuthorityByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorityByIdLogic) GetAuthorityById(in *system.IDReq) (*system.AuthorityResp, error) {
	// todo: add your logic here and delete this line

	return &system.AuthorityResp{}, nil
}
