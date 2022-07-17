package util

import (
	"api/internal/global"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

func TestCaptcha(t *testing.T) {
	driver := base64Captcha.DefaultDriverDigit
	global.GVA_REDIS = redis.New("127.0.0.1:6379")
	var store = NewRedisStore()
	gen := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := gen.Generate()
	fmt.Println(id, b64s, err)
}
