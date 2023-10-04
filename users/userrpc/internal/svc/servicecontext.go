package svc

import (
	"awesomeProject7/users/model"
	"awesomeProject7/users/userrpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	UserModel2 model.UserModel
	Conn1      sqlx.SqlConn
	Conn2      sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Dsn1
	conn := sqlx.NewMysql(dsn)
	dsn2 := c.Dsn2
	conn2 := sqlx.NewMysql(dsn2)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserModel(conn),
		UserModel2: model.NewUserModel(conn2),
		Conn1:      conn,
		Conn2:      conn2,
	}
}
