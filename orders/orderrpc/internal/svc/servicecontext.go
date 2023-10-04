package svc

import (
	"awesomeProject7/orders/model"
	"awesomeProject7/orders/orderrpc/internal/config"
	"awesomeProject7/users/userrpc/userrpcclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrdersModel
	UserRpc    userrpcclient.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Dsn
	conn := sqlx.NewMysql(dsn)
	_ = conn
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrdersModel(conn),
		UserRpc:    userrpcclient.NewUserrpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
