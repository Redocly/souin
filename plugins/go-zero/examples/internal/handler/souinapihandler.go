package handler

import (
	"net/http"

	"github.com/Redocly/souin/plugins/go-zero/examples/internal/logic"
	"github.com/Redocly/souin/plugins/go-zero/examples/internal/svc"
	"github.com/Redocly/souin/plugins/go-zero/examples/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func souin_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CacheReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSouin_apiLogic(r.Context(), svcCtx)
		resp, err := l.Souin_api(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
