package svc

import (
	"awesomeProject7/users/model"
	"awesomeProject7/users/userrpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Dsn
	conn := sqlx.NewMysql(dsn)
	_ = conn
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
