package handler

import (
	"net/http"

	"gas-td-importer/td/internal/svc"
)

func MockDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// 青海
		w.Write([]byte("{\"Status\":true,\"Code\":0,\"Message\":\"查询成功\",\"Data\":[{\"LineName\":\"涩宁兰线\",\"LineCode\":\"10001\",\"Parameters\":[{\"ParaName\":\"压力\",\"ParaCode\":\"YL\",\"ParaValue\":\"2.837335\"},{\"ParaName\":\"流量\",\"ParaCode\":\"LL\",\"ParaValue\":\"464686.27\"}]}]}"))
		// 塔里木
		//w.Write([]byte("\"[{'TagName':'ZTYXC190LNJQZ1_PT9107','TimeStamp':'2021-11-20 08:00:00','Value':6.58229923248291,'Confidence':0,'HostName':'10.79.84.172','Units':'','Tolerance':null},{'TagName':'ZTYXC190LNJQZ1_PT9107','TimeStamp':'2021-11-20 09:00:00','Value':6.58229923248291,'Confidence':0,'HostName':'10.79.84.172','Units':'','Tolerance':null}]\""))
		// 西南
		//w.Write([]byte("[{\"LASTPV_QUALITY\":\"192\",\"STREXT1\":\"CXBWLW/OPC\",\"LASTPV_TIME\":\"2022-06-27 10:01:09\",\"STREXT2\":\"江油中63井.物联网关.A240007PI0HART1401_\",\"UNIT\":\"MPa\",\"STREXT3\":\"中75\",\"STREXT4\":\"0\",\"FLOATEXT4\":\"0.000000\",\"FLOATEXT3\":\"0.000000\",\"FLOATEXT2\":\"0.000000\",\"FLOATEXT1\":\"0.000000\",\"DIGITIALSET\":\"0\",\"TAGTYPE\":\"2\",\"DATATYPE\":\"11\",\"TAGDESC\":\"川西北江油作业区中75井油压（HART）\",\"PV_QUALITY\":\"192\",\"ID\":\"15\",\"tag\":\"/A240007PI0HART1401_\",\"time\":\"2022-06-27 10:01:10\",\"value\":\"6.044363\"},{\"LASTPV_QUALITY\":\"192\",\"STREXT1\":\"CXBWLW/OPC\",\"LASTPV_TIME\":\"2022-06-27 10:01:07\",\"STREXT2\":\"江油中63井.物联网关.A240007PI0HART1402_\",\"UNIT\":\"MPa\",\"STREXT3\":\"中75\",\"STREXT4\":\"0\",\"FLOATEXT4\":\"0.000000\",\"FLOATEXT3\":\"0.000000\",\"FLOATEXT2\":\"0.000000\",\"FLOATEXT1\":\"0.000000\",\"DIGITIALSET\":\"0\",\"TAGTYPE\":\"2\",\"DATATYPE\":\"11\",\"TAGDESC\":\"川西北江油作业区中75井套压（HART）\",\"PV_QUALITY\":\"192\",\"ID\":\"16\",\"tag\":\"/A240007PI0HART1402_\",\"time\":\"2022-06-27 10:01:10\",\"value\":\"6.161860\"}]"))

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
