package handler

import (
	"net/http"

	"cyntra/gateway/api/internal/logic"
	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
