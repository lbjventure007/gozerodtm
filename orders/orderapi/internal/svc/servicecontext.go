package svc

import (
	"awesomeProject7/orders/orderapi/internal/config"
	"awesomeProject7/orders/orderrpc/orderrpcclient"
	"awesomeProject7/users/userrpc/userrpcclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  userrpcclient.Userrpc
	OrderRpc orderrpcclient.Orderrpc
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:   c,
		UserRpc:  userrpcclient.NewUserrpc(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderrpcclient.NewOrderrpc(zrpc.MustNewClient(c.OrderRpc)),
	}
}
