package svc

import (
	"system/rpc/internal/config"
	"system/rpc/internal/initialize"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
	Role   *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.InitGORM(c)
	redis := initialize.InitRedis(c)
	//casbin := initialize.InitCasbin(db)
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
		Role:   nil,
	}
}
