package logic

import (
	"context"
	"fmt"
	"log"

	"gas-td-importer/td/internal/common/errorx"
	"gas-td-importer/td/internal/models"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GasMultiAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGasMultiAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GasMultiAddLogic {
	return &GasMultiAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GasMultiAddLogic) GasMultiAdd(req *types.GasMultiAddRequest) (resp *types.GasMultiAddReply, err error) {
	c := l.svcCtx.Config
	taos := l.svcCtx.TaosEngine

	sql := `INSERT INTO `
	repeat_sql := `%s.%s USING %s.%s (point, pname, unit, region, gases, gas, site, pipeline, uptype, ptype) TAGS('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') VALUES `

	// 遍历gasList 拼接多value insert
	models.PLocker.RLock() // get读锁
	for i := 0; i < len(req.GasList); i++ {
		item := req.GasList[i]

		// 根据point查询PointInfo
		pi, ok := models.PMap[item.Point]
		if !ok {
			log.Printf("找不到%s点位的PointInfo\n", item.Point)
			continue
		}

		// 单位转换
		fmt.Println(item.Value)
		models.ConvertUnit(&pi, &item)
		fmt.Println(item.Value)

		table_sql := fmt.Sprintf(repeat_sql, c.TD.DataBase, item.Point, c.TD.DataBase, c.TD.STable, pi.Point, pi.Pname, pi.Unit,
			pi.Region, pi.Gases, pi.Gas, pi.Site, pi.Pipeline, pi.Uptype, pi.Ptype)
		values_sql := fmt.Sprintf(`('%s', %f) `, item.Ts, item.Value)
		sql += table_sql + values_sql

	}
	models.PLocker.RUnlock()

	fmt.Println("sql:", sql)
	result, err := taos.Exec(sql)
	if err != nil {
		l.Logger.Error("insert failed: " + sql)
		return nil, errorx.NewDefaultError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.GasMultiAddReply{
		Code:    errorx.OKCode,
		Num:     rowsAffected,
		Message: fmt.Sprintf("insert %d lines successed", rowsAffected),
	}, nil
}
