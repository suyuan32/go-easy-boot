package captcha

import (
	"api/internal/util"
	"context"
	"net/http"

	"api/internal/global"
	"api/internal/svc"
	"api/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var store = util.NewRedisStore()

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaInfoResp, err error) {
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth,
		global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	gen := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := gen.Generate(); err != nil {
		logx.Error("getcaptchalogic: fail to generate captcha!", err)
		return nil, err
	} else {
		resp = &types.CaptchaInfoResp{
			BaseMsg: types.BaseMsg{
				Code: http.StatusOK,
				Msg:  "ok",
				Data: "",
			},
			Data: types.CaptchaInfo{
				CaptchaId: id,
				ImgPath:   b64s,
			},
		}
		return resp, nil
	}
}
