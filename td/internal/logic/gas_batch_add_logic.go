package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
