package logic

import (
	"context"
	"system/rpc/internal/model"
	"system/rpc/internal/util"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

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
	result := l.svcCtx.DB.Omit("Role").Create(&model.User{
		UUID:     uuid.New(),
		Username: in.Username,
		Nickname: in.Username,
		Password: util.BcryptEncrypt(in.Password),
		Email:    in.Email,
		Role: model.Role{
			Model: gorm.Model{ID: 2},
		},
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
