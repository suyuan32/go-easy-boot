package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"system/rpc/internal/global"
	"system/rpc/internal/model"
	"system/rpc/internal/util"

	"system/rpc/internal/svc"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *system.RegisterReq) (*system.BaseResp, error) {
	result := global.GVA_DB.Create(&model.User{
		Username: in.UserName,
		Password: util.BcryptEncrypt(in.Password),
		Email:    in.Email,
	})

	if result.Error != nil {
		return &system.BaseResp{
			Code: uint32(codes.Internal),
			Msg:  result.Error.Error(),
		}, result.Error
	}

	return &system.BaseResp{
		Code: uint32(codes.OK),
		Msg:  "successful",
	}, nil
}
