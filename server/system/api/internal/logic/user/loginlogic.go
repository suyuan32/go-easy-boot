package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"net/http"
	"system/rpc/types/system"

	"system/api/internal/svc"
	"system/api/internal/types"

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
	if ok := store.Verify(req.CaptchaId, req.Captcha, true); ok {

		user, err := l.svcCtx.SystemRpc.Login(context.Background(),
			&system.LoginReq{
				Username: req.Username,
				Password: req.Password,
			})
		if err != nil {
			return nil, err
		}
		if user.Code == uint32(codes.OK) {
			token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, 100,
				l.svcCtx.Config.Auth.AccessExpire, int64(user.Id))
			if err != nil {
				return nil, err
			}
			resp = &types.LoginResp{
				BaseMsg: types.BaseMsg{Code: http.StatusOK, Msg: "login successfully"},
				Data: types.LoginRespData{
					UserId:       user.Id,
					Username:     user.Username,
					Avatar:       user.Avatar,
					RoleId:       user.AuthorityId,
					AccessToken:  token,
					AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
				},
			}
			return resp, nil
		} else if user.Code == uint32(codes.NotFound) {
			resp = &types.LoginResp{
				BaseMsg: types.BaseMsg{Code: http.StatusNotFound, Msg: "user does not exist"},
				Data:    types.LoginRespData{},
			}
			return resp, nil
		} else if user.Code == uint32(codes.Internal) {
			resp = &types.LoginResp{
				BaseMsg: types.BaseMsg{Code: http.StatusNotFound, Msg: "user does not exist"},
				Data:    types.LoginRespData{},
			}
			return resp, nil
		}
	} else {
		resp = &types.LoginResp{
			BaseMsg: types.BaseMsg{Code: http.StatusNotFound, Msg: "wrong captcha"},
			Data:    types.LoginRespData{},
		}
		return resp, nil
	}
	return nil, err
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
