package user

import (
	"net/http"

	"api/internal/logic/user"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewModifyInfoLogic(r.Context(), svcCtx)
		resp, err := l.ModifyInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
