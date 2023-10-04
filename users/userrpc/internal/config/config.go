package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	//UserModel model.UserModel
	Dsn1   string
	Dsn2   string
	Consul consul.Conf
}
