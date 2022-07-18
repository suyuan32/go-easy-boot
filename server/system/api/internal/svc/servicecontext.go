package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"system/api/internal/config"
	"system/api/internal/middleware"
	"system/rpc/system"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Cros      rest.Middleware
	SystemRpc system.System
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Cros:      middleware.NewCrosMiddleware().Handle,
		SystemRpc: system.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
	}
}
