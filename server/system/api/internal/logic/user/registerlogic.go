package user

import (
	"context"
	"google.golang.org/grpc/codes"
	"net/http"
	"system/api/internal/util"
	"system/rpc/types/system"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var store = util.NewRedisStore()

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if ok := store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.SystemRpc.CreateUser(context.Background(),
			&system.RegisterReq{
				Username: req.Username,
				Password: req.Password,
				Email:    req.Email,
			})
		if err != nil {
			l.Logger.Error("register logic: create user err: ", err.Error())
			return nil, err
		}
		if user.Code == uint32(codes.OK) {
			resp = &types.RegisterResp{
				BaseMsg: types.BaseMsg{Code: http.StatusOK, Msg: "ok"},
				Data:    types.RegisterRespData{Status: "ok"},
			}
			return resp, nil
		} else if user.Code == uint32(codes.AlreadyExists) {
			resp = &types.RegisterResp{
				BaseMsg: types.BaseMsg{Code: http.StatusOK, Msg: "exist"},
				Data:    types.RegisterRespData{Status: "fail"},
			}
			return resp, nil
		} else if user.Code == uint32(codes.Internal) {
			resp = &types.RegisterResp{
				BaseMsg: types.BaseMsg{Code: http.StatusInternalServerError, Msg: "internal error"},
				Data:    types.RegisterRespData{Status: "fail"},
			}
			return resp, nil
		}
	} else {
		resp = &types.RegisterResp{
			BaseMsg: types.BaseMsg{Code: http.StatusBadRequest, Msg: "wrong captcha"},
			Data:    types.RegisterRespData{Status: "fail"},
		}
		return resp, nil
	}
	return nil, err
}
