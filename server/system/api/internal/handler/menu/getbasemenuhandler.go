package menu

import (
	"net/http"

	"api/internal/logic/menu"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBaseMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewGetBaseMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetBaseMenu(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
