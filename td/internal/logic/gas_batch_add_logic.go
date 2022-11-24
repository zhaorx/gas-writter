package logic

import (
	"context"
	"fmt"

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
	resp = new(types.GasBatchAddReply)
	insert_sql := `INSERT INTO %s.%s USING %s.%s TAGS('%s', '%s', '%s', '%s') VALUES `

	// 拼接多value insert
	suffix := ""
	for i := 0; i < len(req.Tss); i++ {
		suffix += fmt.Sprintf(`('%s', %f)`, req.Tss[i], req.Values[i])
	}
	sql := fmt.Sprintf(insert_sql, c.TD.DataBase, req.Point, c.TD.DataBase, c.TD.STable, req.Point, req.PName, req.Unit, req.Region) + suffix
	result, err := taos.Exec(sql)
	if err != nil {
		l.Logger.Error("insert failed: " + sql)
		fmt.Println("failed to insert, err:", err)
		resp.Message = fmt.Sprintf("insert failed:%s", err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	resp.Num = rowsAffected
	if err != nil {
		fmt.Println("failed to get affected rows, err:", err)
	}

	if rowsAffected >= 1 {
		resp.Message = fmt.Sprintf("insert %d lines successed", rowsAffected)
	} else {
		resp.Message = fmt.Sprintf("insert failed:%s", err.Error())
	}

	return
}
