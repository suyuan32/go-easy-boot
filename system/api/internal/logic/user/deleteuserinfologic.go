package user

import (
	"context"

	"system/api/internal/svc"
	"system/api/internal/types"
	"system/rpc/types/system"

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
	_, err = l.svcCtx.SystemRpc.DeleteRole(context.Background(), &system.IDReq{
		ID: 0,
	})
	return nil, err
}
