package user

import (
	"context"
	"net/http"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	resp = &types.LoginResp{
		BaseMsg: types.BaseMsg{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: "",
		},
		Data: types.LoginRespData{
			UserId:       1,
			Username:     "jack",
			Avatar:       "/haha",
			RoleId:       1,
			AccessToken:  "dsafasdfasfasdfasdf",
			AccessExpire: 100000,
		},
	}

	return resp, nil
}
