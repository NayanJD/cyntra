package handler

import (
	"net/http"

	"cyntra/common/errorx"
	"cyntra/gateway/api/internal/logic"
	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type GenericResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewValidationCodeResponse(err))
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, errorx.NewDefaultSuccess(resp))
		}
	}
}
