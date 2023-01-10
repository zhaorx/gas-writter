package logic

import (
	"context"
	"fmt"
	"log"

	"gas-td-importer/td/internal/common/errorx"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GasListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGasListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GasListLogic {
	return &GasListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GasListLogic) GasList(req *types.GasListRequest) (resp *types.GasListReply, err error) {
	c := l.svcCtx.Config
	taos := l.svcCtx.TaosEngine

	query_sql := `SELECT * FROM %s.%s where ts >= '%s' AND ts <= '%s'`
	sql := fmt.Sprintf(query_sql, c.TD.DataBase, c.TD.STable, req.TsStart, req.TsEnd)

	rows, err := taos.Query(sql)
	if err != nil {
		log.Fatalln("failed to select from table, err:", err)
	}
	defer rows.Close()
	list := make([]*types.Gas, 0)
	for rows.Next() {
		var r = types.Gas{
			Ts:     "",
			Value:  0,
			Point:  "",
			PName:  "",
			Unit:   "",
			Region: "",
		}
		err := rows.Scan(&r.Ts, &r.Value, &r.Point, &r.PName, &r.Unit, &r.Region)
		if err != nil {
			log.Fatalln("scan error:\n", err)
		}

		list = append(list, &r)
	}

	return &types.GasListReply{
		Code:    errorx.OKCode,
		Message: fmt.Sprintf("select successed"),
		List:    list,
	}, nil
}
