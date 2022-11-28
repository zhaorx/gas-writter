package handler

import (
	"net/http"

	"gas-td-importer/td/internal/logic"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GasListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GasListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if start := r.Form.Get("ts_start"); len(start) > 0 && len(req.TsStart) == 0 {
			req.TsStart = start
		}
		if end := r.Form.Get("ts_end"); len(end) > 0 && len(req.TsEnd) == 0 {
			req.TsEnd = end
		}

		l := logic.NewGasListLogic(r.Context(), svcCtx)
		resp, err := l.GasList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
