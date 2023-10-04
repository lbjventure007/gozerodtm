package main

import (
	//_ "github.com/dtm-labs/driver-gozero"

	"awesomeProject7/orders/orderapi/internal/config"
	"awesomeProject7/orders/orderapi/internal/handler"
	"awesomeProject7/orders/orderapi/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	//"github.com/zeromicro/zero-contrib/rest/registry/etcd"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "etc/orderapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//dtmdriver.Register
	server := rest.MustNewServer(c.RestConf)
	//	dtmdriver.Register("")
	//err := dtmdriver.Use("dtmservice")
	//fmt.Println("----", err)
	//etcd.RegisterRest(c.Etcd, c.RestConf)

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
