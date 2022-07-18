package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/api/internal/logic/user"
	"system/api/internal/svc"
	"system/api/internal/types"
)

func GetUserInfoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserInfoListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
