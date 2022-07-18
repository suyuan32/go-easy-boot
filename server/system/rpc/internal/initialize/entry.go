package initialize

import (
	"github.com/zeromicro/go-zero/core/logx"
	"system/rpc/internal/global"
)

func InitAll() {
	// initialize database
	global.GVA_DB = InitGORM()
	if global.GVA_DB == nil {
		logx.Error("fail to initialize database")
	}

	// initialize role access control
	InitCasbin()
}
