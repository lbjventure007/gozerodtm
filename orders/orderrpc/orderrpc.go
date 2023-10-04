package main

import (
	"awesomeProject7/orders/orderrpc/internal/config"
	"awesomeProject7/orders/orderrpc/internal/server"
	"awesomeProject7/orders/orderrpc/internal/svc"
	"awesomeProject7/orders/orderrpc/orderrpc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/orderrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		orderrpc.RegisterOrderrpcServer(grpcServer, server.NewOrderrpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	//fmt.Println("-----driver----", err)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
