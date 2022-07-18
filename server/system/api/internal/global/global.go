package global

import (
	"system/api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	GVA_CONFIG *config.Config
	GVA_REDIS  *redis.Redis
)
