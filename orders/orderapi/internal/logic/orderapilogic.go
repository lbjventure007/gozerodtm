package logic

import (
	"context"

	"awesomeProject7/orders/orderapi/internal/svc"
	"awesomeProject7/orders/orderapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderapiLogic {
	return &OrderapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderapiLogic) Orderapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
