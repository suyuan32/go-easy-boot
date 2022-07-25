package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"
	"system/api/internal/types"

	"system/api/common/errorx"
	"system/api/internal/config"
	"system/api/internal/handler"
	"system/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "./etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	initErrorHandler()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// initialize the error handler
func initErrorHandler() {
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.ApiError:
			switch e.Code {
			case 0:
				return http.StatusOK, &types.SimpleMsg{
					Msg: status.Convert(err).Message(),
				}
			default:
				return e.Code, &types.SimpleMsg{
					Msg: status.Convert(err).Message(),
				}
			}
		default:
			return errorx.CodeFromGrpcError(e), &types.SimpleMsg{
				Msg: status.Convert(err).Message(),
			}
		}
	})
}
