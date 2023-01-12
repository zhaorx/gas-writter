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

type GasAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGasAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GasAddLogic {
	return &GasAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GasAddLogic) GasAdd(req *types.GasAddRequest) (resp *types.GasAddReply, err error) {
	c := l.svcCtx.Config
	taos := l.svcCtx.TaosEngine

	gas := types.Gas{
		Ts:    req.Ts,
		Value: req.Value,
		Point: req.Point,
	}

	// 根据point查询PointInfo
	models.PLocker.RLock() // get读锁
	pi, ok := models.PMap[gas.Point]
	models.PLocker.RUnlock()
	if !ok {
		err := errorx.NewDefaultError(fmt.Sprintf("找不到%s点位的PointInfo", gas.Point))
		log.Println(err.Error())
		return &types.GasAddReply{
			Code:    errorx.DefaultErrorCode,
			Num:     0,
			Message: err.Error(),
		}, nil
	}

	// 单位转换
	models.ConvertUnit(&pi, &gas)

	insert_sql := `
		INSERT INTO %s.%s USING %s.%s (point, pname, unit, region, gases, gas, site, pipeline, uptype, ptype)
        TAGS('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') VALUES ('%s', %f)
	`
	sql := fmt.Sprintf(insert_sql, c.TD.DataBase, gas.Point, c.TD.DataBase, c.TD.STable,
		pi.Point, pi.Pname, pi.Unit, pi.Region, pi.Gases, pi.Gas, pi.Site, pi.Pipeline, pi.Uptype, pi.Ptype,
		gas.Ts, gas.Value)
	result, err := taos.Exec(sql)
	if err != nil {
		l.Logger.Error("insert failed: " + sql)
		return nil, errorx.NewDefaultError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.GasAddReply{
		Code:    errorx.OKCode,
		Num:     rowsAffected,
		Message: fmt.Sprintf("insert %d line successed", rowsAffected),
	}, nil
}
