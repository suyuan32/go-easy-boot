package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"system/rpc/internal/model"
	"system/rpc/internal/svc"
	"system/rpc/internal/util"
	"system/rpc/types/system"

	"github.com/google/uuid"
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
	result := l.svcCtx.DB.Create(&model.User{
		UUID:     uuid.NewString(),
		Username: in.Username,
		Nickname: in.Username,
		Password: util.BcryptEncrypt(in.Password),
		Email:    in.Email,
		RoleId:   2,
	})

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &system.BaseResp{
		Msg: "successful",
	}, nil
}
