package svc

import (
	"api/internal/config"
	"api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Cros   rest.Middleware
	//SystemRpc system.System
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cros:   middleware.NewCrosMiddleware().Handle,
	}
}
