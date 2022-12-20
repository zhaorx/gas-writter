package logic

import (
	"context"
	"fmt"

	"gas-td-importer/td/internal/common/errorx"
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
	taos := l.svcCtx.Engine

	sql := `INSERT INTO `
	repeat_sql := `%s.%s USING %s.%s TAGS('%s', '%s', '%s', '%s') VALUES `

	// 遍历gasList 拼接多value insert
	for i := 0; i < len(req.GasList); i++ {
		item := req.GasList[i]

		table_sql := fmt.Sprintf(repeat_sql, c.TD.DataBase, item.Point, c.TD.DataBase, c.TD.STable, item.Point, item.PName, item.Unit, item.Region)
		values_sql := fmt.Sprintf(`('%s', %f) `, item.Ts, item.Value)
		sql += table_sql + values_sql

	}

	//fmt.Println("sql:", sql)
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

	return
}
