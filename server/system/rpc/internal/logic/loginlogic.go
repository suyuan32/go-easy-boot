package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"system/rpc/internal/model"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  user service
func (l *LoginLogic) Login(in *system.LoginReq) (*system.LoginResp, error) {
	var u model.User
	result := l.svcCtx.DB.Where(&model.User{Username: in.Username}).First(&u)
	if result.Error != nil {
		l.Logger.Error("login logic: database error ", result.Error)
		return &system.LoginResp{
			Code:        uint32(codes.Internal),
			Id:          0,
			AuthorityId: 0,
			Status:      0,
			Username:    "",
			Avatar:      "",
		}, result.Error
	}
	if result.RowsAffected == 0 {
		return &system.LoginResp{
			Code:        uint32(codes.NotFound),
			Id:          0,
			AuthorityId: 0,
			Status:      0,
			Username:    "",
			Avatar:      "",
		}, nil
	} else {
		return &system.LoginResp{
			Code:        uint32(codes.OK),
			Id:          uint64(u.ID),
			AuthorityId: int32(u.AuthorityId),
			Status:      u.Status,
			Username:    u.Username,
			Avatar:      u.Avatar,
		}, nil
	}
}
