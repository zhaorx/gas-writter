package handler

import (
	"net/http"

	"gas-td-importer/td/internal/logic"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GasAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GasAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGasAddLogic(r.Context(), svcCtx)
		resp, err := l.GasAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
