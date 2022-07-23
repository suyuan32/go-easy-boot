package logic

import (
	"context"
	"fmt"

	"system/rpc/internal/model"
	"system/rpc/internal/svc"
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
		return &system.LoginResp{
			Code: uint32(codes.Internal),
		}, nil
	}

	if result.RowsAffected == 0 {
		return &system.LoginResp{
			Code: uint32(codes.NotFound),
		}, nil
	}

	// get role data from redis
	var roleName string
	if s, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleId)); err != nil {
		var roleData []model.Role
		res := l.svcCtx.DB.Find(&roleData)
		if res.RowsAffected == 0 {
			return &system.LoginResp{
				Code: uint32(codes.Internal),
			}, nil
		}
		for _, v := range roleData {
			err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", u.RoleId), v.Name)
			if err != nil {
				return &system.LoginResp{
					Code: uint32(codes.Internal),
				}, nil
			}
			if v.RoleId == uint64(u.RoleId) {
				roleName = v.Name
			}
		}
	} else {
		roleName = s
	}

	return &system.LoginResp{
		Code:     uint32(codes.OK),
		Id:       u.UUID,
		RoleId:   u.RoleId,
		RoleName: roleName,
	}, nil

}
