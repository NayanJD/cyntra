package handler

import (
	"net/http"

	"cyntra/common/errorx"
	"cyntra/gateway/api/internal/logic"
	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewCodeResponse(codes.InvalidArgument, codes.InvalidArgument.String(), err.Error()))
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
