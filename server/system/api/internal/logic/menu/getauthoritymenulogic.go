package menu

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorityMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorityMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityMenuLogic {
	return &GetAuthorityMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorityMenuLogic) GetAuthorityMenu(req *types.IdReq) (resp *types.AuthorityMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
