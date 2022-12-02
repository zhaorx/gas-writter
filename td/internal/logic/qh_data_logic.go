package logic

import (
	"context"

	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QhDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQhDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QhDataLogic {
	return &QhDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QhDataLogic) QhData(req *types.QhDataRequest) (resp *types.QhDataReply, err error) {
	// todo: add your logic here and delete this line

	return
}
