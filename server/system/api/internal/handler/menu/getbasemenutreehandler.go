package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/api/internal/logic/menu"
	"system/api/internal/svc"
)

func GetBaseMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetBaseMenuTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetBaseMenuTree()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
