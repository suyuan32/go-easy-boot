package user

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoListLogic {
	return &GetUserInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoListLogic) GetUserInfoList(req *types.PageInfo) (resp *types.UserInfoListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
