package global

import (
	"system/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
	GVA_REDIS  *redis.Redis
)
