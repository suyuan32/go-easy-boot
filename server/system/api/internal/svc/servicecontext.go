package svc

import (
	"system/api/internal/config"
	"system/api/internal/initialize"
	"system/api/internal/middleware"
	"system/rpc/system"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Cros      rest.Middleware
	SystemRpc system.System
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := initialize.InitRedis(c)
	return &ServiceContext{
		Config:    c,
		Cros:      middleware.NewCrosMiddleware().Handle,
		SystemRpc: system.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
		Redis:     redis,
	}
}
