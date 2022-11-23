package handler

import (
	"net/http"

	"gas-td-importer/td/internal/logic"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GasBatchAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GasBatchAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGasBatchAddLogic(r.Context(), svcCtx)
		resp, err := l.GasBatchAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
