// Code generated by goctl. DO NOT EDIT.
// Source: orderrpc.proto

package server

import (
	"context"

	"awesomeProject7/orders/orderrpc/internal/logic"
	"awesomeProject7/orders/orderrpc/internal/svc"
	"awesomeProject7/orders/orderrpc/orderrpc"
)

type OrderrpcServer struct {
	svcCtx *svc.ServiceContext
	orderrpc.UnimplementedOrderrpcServer
}

func NewOrderrpcServer(svcCtx *svc.ServiceContext) *OrderrpcServer {
	return &OrderrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderrpcServer) Ping(ctx context.Context, in *orderrpc.Request) (*orderrpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *OrderrpcServer) CreateOrder(ctx context.Context, in *orderrpc.CreateOrderRequest) (*orderrpc.CreateOrderResponse, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}
