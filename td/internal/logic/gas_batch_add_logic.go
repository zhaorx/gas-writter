package logic

import (
	"context"
	"fmt"

	"gas-td-importer/td/internal/common/errorx"
	"gas-td-importer/td/internal/models"
	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GasBatchAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGasBatchAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GasBatchAddLogic {
	return &GasBatchAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GasBatchAddLogic) GasBatchAdd(req *types.GasBatchAddRequest) (resp *types.GasBatchAddReply, err error) {
	c := l.svcCtx.Config
	taos := l.svcCtx.TaosEngine

	// 根据point查询PointInfo
	models.PLocker.RLock() // get读锁
	pi, ok := models.PMap[req.Point]
	models.PLocker.RUnlock()
	if !ok {
		err := errorx.NewDefaultError(fmt.Sprintf("找不到%s点位的PointInfo", req.Point))
		l.Logger.Errorf(err.Error())
		return &types.GasBatchAddReply{
			Code:    errorx.DefaultErrorCode,
			Num:     0,
			Message: err.Error(),
		}, nil
	}

	// 拼接多value insert
	suffix := ""
	for i := 0; i < len(req.Tss); i++ {
		gas := types.Gas{
			Ts:    req.Tss[i],
			Value: req.Values[i],
			Point: req.Point,
		}

		// 单位转换
		models.ConvertUnit(&pi, &gas)
		suffix += fmt.Sprintf(`('%s', %f)`, gas.Ts, gas.Value)
	}

	insert_sql := `INSERT INTO %s.%s USING %s.%s (point, pname, unit, region, gases, gas, site, pipeline, uptype, ptype) TAGS('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') VALUES `
	sql := fmt.Sprintf(insert_sql, c.TD.DataBase, req.Point, c.TD.DataBase, c.TD.STable, pi.Point, pi.Pname, pi.Unit, pi.Region, pi.Gases, pi.Gas, pi.Site, pi.Pipeline, pi.Uptype, pi.Ptype) + suffix
	result, err := taos.Exec(sql)
	if err != nil {
		l.Logger.Error("insert failed: " + sql)
		return nil, errorx.NewDefaultError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.GasBatchAddReply{
		Code:    errorx.OKCode,
		Num:     rowsAffected,
		Message: fmt.Sprintf("insert %d lines successed", rowsAffected),
	}, nil
}
