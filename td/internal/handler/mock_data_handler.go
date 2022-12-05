package handler

import (
	"net/http"

	"gas-td-importer/td/internal/svc"
)

func MockDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 青海
		//w.Header().Set("Content-Type", "application/json; charset=utf-8")
		//w.Write([]byte("{\"Status\":true,\"Code\":0,\"Message\":\"查询成功\",\"Data\":[{\"LineName\":\"涩宁兰线\",\"LineCode\":\"10001\",\"Parameters\":[{\"ParaName\":\"压力\",\"ParaCode\":\"YL\",\"ParaValue\":\"2.837335\"},{\"ParaName\":\"流量\",\"ParaCode\":\"LL\",\"ParaValue\":\"464686.27\"}]}]}"))

		// 塔里木
		//w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("\"[{'TagName':'ZTYXC190LNJQZ1_PT9107','TimeStamp':'2021-11-20 08:00:00','Value':6.58229923248291,'Confidence':0,'HostName':'10.79.84.172','Units':'','Tolerance':null},{'TagName':'ZTYXC190LNJQZ1_PT9107','TimeStamp':'2021-11-20 09:00:00','Value':6.58229923248291,'Confidence':0,'HostName':'10.79.84.172','Units':'','Tolerance':null}]\""))

		//var req types.MockDataRequest
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		//
		//l := logic.NewMockDataLogic(r.Context(), svcCtx)
		//resp, err := l.MockData(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
