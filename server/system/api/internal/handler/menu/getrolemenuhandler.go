package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/api/internal/logic/menu"
	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRoleMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewGetRoleMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleMenu(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
