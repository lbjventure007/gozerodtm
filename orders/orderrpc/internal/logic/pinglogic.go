package logic

import (
	"context"

	"awesomeProject7/orders/orderrpc/internal/svc"
	"awesomeProject7/orders/orderrpc/orderrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *orderrpc.Request) (*orderrpc.Response, error) {
	// todo: add your logic here and delete this line

	return &orderrpc.Response{}, nil
}
