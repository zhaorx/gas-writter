package logic

import (
	"context"
	"fmt"

	"gas-td-importer/td/internal/common/errorx"
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
	taos := l.svcCtx.Engine
	insert_sql := `INSERT INTO %s.%s USING %s.%s (point, pname, unit, region) TAGS('%s', '%s', '%s', '%s') VALUES `

	// 拼接多value insert
	suffix := ""
	for i := 0; i < len(req.Tss); i++ {
		suffix += fmt.Sprintf(`('%s', %f)`, req.Tss[i], req.Values[i])
	}
	sql := fmt.Sprintf(insert_sql, c.TD.DataBase, req.Point, c.TD.DataBase, c.TD.STable, req.Point, req.PName, req.Unit, req.Region) + suffix
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
