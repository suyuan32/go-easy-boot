package system

import (
	"net/http"

	"api/internal/logic/system"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewHealthCheckLogic(r.Context(), svcCtx)
		err := l.HealthCheck()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
