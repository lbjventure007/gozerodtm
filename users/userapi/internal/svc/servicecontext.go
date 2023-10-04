package svc

import (
	"awesomeProject7/users/userapi/internal/config"
	"awesomeProject7/users/userrpc/userrpcclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpcclient.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpcclient.NewUserrpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
