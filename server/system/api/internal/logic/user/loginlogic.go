package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"system/api/common/errorx"
	"system/api/internal/logic/captcha"
	"system/api/internal/svc"
	"system/api/internal/types"
	"system/rpc/types/system"

	"github.com/golang-jwt/jwt/v4"
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
	// return nil, errorx.NewRpcError(codes.InvalidArgument, "login request")
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.SystemRpc.Login(context.Background(),
			&system.LoginReq{
				Username: req.Username,
				Password: req.Password,
			})
		// fmt.Println(err.Error())
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, user.Id, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, int64(user.RoleId))
		if err != nil {
			return nil, err
		}
		resp = &types.LoginResp{
			BaseMsg: types.BaseMsg{Code: http.StatusOK, Msg: ""},
			Data: types.LoginRespData{
				UserId: user.Id,
				Token:  token,
				Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
				Role: types.RoleInfo{
					Value:    user.RoleId,
					RoleName: user.RoleName,
				},
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewApiError(http.StatusBadRequest, "sys.login.wrongCaptcha")
	}
}

func (l *LoginLogic) getJwtToken(secretKey, uuid string, iat, seconds, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = uuid
	claims["roleId"] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
