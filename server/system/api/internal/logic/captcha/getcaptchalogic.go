package captcha

import (
	"context"
	"net/http"

	"system/api/common/errorx"
	"system/api/internal/config"
	"system/api/internal/svc"
	"system/api/internal/types"
	"system/api/internal/util"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var Store *util.RedisStore
var driver *base64Captcha.DriverDigit

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaInfoResp, err error) {
	if driver == nil || Store == nil {
		initStoreAndDriver(l.svcCtx.Config, l.svcCtx.Redis)
	}
	gen := base64Captcha.NewCaptcha(driver, Store)
	if id, b64s, err := gen.Generate(); err != nil {
		l.Logger.Error("getcaptchalogic: fail to generate captcha!", err)
		return nil, errorx.NewApiError(http.StatusInternalServerError, "内部错误")
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

func initStoreAndDriver(c config.Config, r *redis.Redis) {
	driver = base64Captcha.NewDriverDigit(c.Captcha.ImgHeight, c.Captcha.ImgWidth,
		c.Captcha.KeyLong, 0.7, 80)
	Store = util.NewRedisStore(r)
}
