package logic

import (
	"context"

	"gas-td-importer/td/internal/svc"
	"gas-td-importer/td/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MockDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMockDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MockDataLogic {
	return &MockDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MockDataLogic) MockData(req *types.MockDataRequest) (resp *types.MockDataReply, err error) {
	// todo: add your logic here and delete this line

	return
}
