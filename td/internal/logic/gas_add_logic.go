package logic

import (
	"context"
	"fmt"

	"gas-td-importer/td/internal/common/errorx"
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
	taos := l.svcCtx.Engine

	insert_sql := `INSERT INTO %s.%s USING %s.%s TAGS('%s', '%s', '%s', '%s') VALUES ('%s', %f)`
	sql := fmt.Sprintf(insert_sql, c.TD.DataBase, req.Point, c.TD.DataBase, c.TD.STable, req.Point, req.PName, req.Unit, req.Region, req.Ts, req.Value)
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
		Message: fmt.Sprintf("insert %d lines successed", rowsAffected),
	}, nil
}
