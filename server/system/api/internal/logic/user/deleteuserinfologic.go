package user

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserInfoLogic {
	return &DeleteUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserInfoLogic) DeleteUserInfo(req *types.IdReq) (resp *types.BaseMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
