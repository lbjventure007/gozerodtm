// Code generated by goctl. DO NOT EDIT.
// Source: userrpc.proto

package userrpcclient

import (
	"context"

	"awesomeProject7/users/userrpc/userrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request               = userrpc.Request
	Response              = userrpc.Response
	UpdateBalanceRequest  = userrpc.UpdateBalanceRequest
	UpdateBalanceResponse = userrpc.UpdateBalanceResponse

	Userrpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error)
	}

	defaultUserrpc struct {
		cli zrpc.Client
	}
)

func NewUserrpc(cli zrpc.Client) Userrpc {
	return &defaultUserrpc{
		cli: cli,
	}
}

func (m *defaultUserrpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := userrpc.NewUserrpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUserrpc) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error) {
	client := userrpc.NewUserrpcClient(m.cli.Conn())
	return client.UpdateBalance(ctx, in, opts...)
}
