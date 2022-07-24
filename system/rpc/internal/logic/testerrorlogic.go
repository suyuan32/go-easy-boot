package logic

import (
	"context"

	"system/api/common/errorx"
	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type TestErrorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestErrorLogic {
	return &TestErrorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestErrorLogic) TestError(in *system.IDReq) (*system.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &system.BaseResp{}, errorx.NewRpcError(codes.Unavailable, "unavailable")
}
