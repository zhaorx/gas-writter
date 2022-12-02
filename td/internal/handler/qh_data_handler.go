package handler

import (
	"net/http"

	"gas-td-importer/td/internal/svc"
)

func QhDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.QhDataRequest

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("{\"Status\":true,\"Code\":0,\"Message\":\"查询成功\",\"Data\":[{\"LineName\":\"涩宁兰线\",\"LineCode\":\"10001\",\"Parameters\":[{\"ParaName\":\"压力\",\"ParaCode\":\"YL\",\"ParaValue\":\"2.837335\"},{\"ParaName\":\"流量\",\"ParaCode\":\"LL\",\"ParaValue\":\"464686.27\"}]}]}"))
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		//
		//l := logic.NewQhDataLogic(r.Context(), svcCtx)
		//resp, err := l.QhData(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
