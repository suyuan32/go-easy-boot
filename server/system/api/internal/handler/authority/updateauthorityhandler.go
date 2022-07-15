package authority

import (
	"net/http"

	"api/internal/logic/authority"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Authority
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewUpdateAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
