package menu

import (
	"net/http"

	"api/internal/logic/menu"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
