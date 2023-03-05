package logic

import (
	"context"

	"hellp/internal/svc"
	"hellp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HellpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHellpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HellpLogic {
	return &HellpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HellpLogic) Hellp(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
