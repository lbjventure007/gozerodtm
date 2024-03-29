// Code generated by goctl. DO NOT EDIT.
// Source: orderrpc.proto

package orderrpcclient

import (
	"context"

	"awesomeProject7/orders/orderrpc/orderrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderRequest  = orderrpc.CreateOrderRequest
	CreateOrderResponse = orderrpc.CreateOrderResponse
	Request             = orderrpc.Request
	Response            = orderrpc.Response

	Orderrpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	}

	defaultOrderrpc struct {
		cli zrpc.Client
	}
)

func NewOrderrpc(cli zrpc.Client) Orderrpc {
	return &defaultOrderrpc{
		cli: cli,
	}
}

func (m *defaultOrderrpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := orderrpc.NewOrderrpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultOrderrpc) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := orderrpc.NewOrderrpcClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}
