package initialize

import (
	"api/internal/global"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func InitRedis() {
	r := redis.New(global.GVA_CONFIG.RedisConf.Host, func(r *redis.Redis) {
		r.Type = global.GVA_CONFIG.RedisConf.Type
		r.Pass = global.GVA_CONFIG.RedisConf.Pass
	})
	global.GVA_REDIS = r
}
