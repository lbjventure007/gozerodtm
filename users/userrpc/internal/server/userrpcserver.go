// Code generated by goctl. DO NOT EDIT.
// Source: userrpc.proto

package server

import (
	"context"

	"awesomeProject7/users/userrpc/internal/logic"
	"awesomeProject7/users/userrpc/internal/svc"
	"awesomeProject7/users/userrpc/userrpc"
)

type UserrpcServer struct {
	svcCtx *svc.ServiceContext
	userrpc.UnimplementedUserrpcServer
}

func NewUserrpcServer(svcCtx *svc.ServiceContext) *UserrpcServer {
	return &UserrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserrpcServer) Ping(ctx context.Context, in *userrpc.Request) (*userrpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserrpcServer) UpdateBalance(ctx context.Context, in *userrpc.UpdateBalanceRequest) (*userrpc.UpdateBalanceResponse, error) {
	l := logic.NewUpdateBalanceLogic(ctx, s.svcCtx)
	return l.UpdateBalance(in)
}