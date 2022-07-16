package global

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"rpc/internal/config"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
	GVA_REDIS  *redis.Redis
)
