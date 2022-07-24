package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"

	"system/api/common/errorx"
	"system/rpc/internal/model"
	"system/rpc/internal/svc"
	"system/rpc/internal/util"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
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
		return nil, status.Error(codes.Internal, "database err")
	}

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "user not exist")
	}

	if ok := util.BcryptCheck(in.Password, u.Password); !ok {
		return nil, status.Error(codes.InvalidArgument, "account or password incorrect")
	}

	// get role data from redis
	var roleName string
	if s, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId)); err != nil {
		var roleData []model.Role
		res := l.svcCtx.DB.Find(&roleData)
		if res.RowsAffected == 0 {
			return nil, errorx.NewRpcError(codes.NotFound, "role not exist")
		}
		for _, v := range roleData {
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", u.RoleId), v.Name)
			if err != nil {
				return nil, errorx.NewRpcError(codes.Internal, "redis error")
			}
			if v.RoleId == uint64(u.RoleId) {
				roleName = v.Name
			}
		}
	} else {
		roleName = s
	}

	return &system.LoginResp{
		Id:       u.UUID,
		RoleId:   u.RoleId,
		RoleName: roleName,
	}, nil

}
