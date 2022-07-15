package global

import (
	"gorm.io/gorm"
	"rpc/internal/config"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
)
