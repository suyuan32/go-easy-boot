package main

import (
	"flag"
	"fmt"

	"system/api/internal/config"
	"system/api/internal/global"
	"system/api/internal/handler"
	"system/api/internal/initialize"
	"system/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	global.GVA_CONFIG = &c
	initialize.InitAll()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
