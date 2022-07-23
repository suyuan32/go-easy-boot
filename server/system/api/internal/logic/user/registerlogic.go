package user

import (
	"context"
	"net/http"

	"system/api/internal/logic/captcha"
	"system/api/internal/svc"
	"system/api/internal/types"
	"system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResp, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
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
			resp = &types.BaseResp{
				Code: http.StatusOK,
				Msg:  "ok",
			}
			return resp, nil
		} else if user.Code == uint32(codes.AlreadyExists) {
			resp = &types.BaseResp{
				Code: http.StatusBadRequest,
				Msg:  "sys.login.signupUserExist",
			}
			return resp, nil
		} else if user.Code == uint32(codes.Internal) {
			resp = &types.BaseResp{
				Code: http.StatusInternalServerError,
				Msg:  "internal error",
			}
			return resp, nil
		}
	} else {
		resp = &types.BaseResp{
			Code: http.StatusBadRequest,
			Msg:  "wrong captcha",
		}
		return resp, nil
	}
	return nil, err
}
