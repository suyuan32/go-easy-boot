package system

import (
	"context"

	"system/api/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestErrorLogic {
	return &TestErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestErrorLogic) TestError() error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.SystemRpc.TestError(context.Background(), &system.IDReq{
		ID: 0,
	})
	return err
}
