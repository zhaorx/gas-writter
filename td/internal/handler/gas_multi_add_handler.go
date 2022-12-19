package handler

import (
	"net/http"

	"gas-td-importer/td/internal/logic"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GasMultiAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GasMultiAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGasMultiAddLogic(r.Context(), svcCtx)
		resp, err := l.GasMultiAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
