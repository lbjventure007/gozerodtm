# 如何生成api服务
可以使用goctl命令生成必要的文件和目录

goctl api go -api *.api -dir ./
goctl api go 表示生成go语言的服务
api *.api 指定api文件
go-zero 定义的api文件 重新生成对应的go代码使用  表示当前test.api 对应的生成到当前根目录下
goctl api go -api test.api -dir .     
如果有多个api文件 可以使用
goctl api go -api *.api -dir .

# 如何生成rpc服务
# 单个 rpc 服务生成示例指令
$ goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
# 多个 rpc 服务生成示例指令
$ goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. -m


go-zero 定义的proto 后  生成对于到go文件命令： 
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# 所需组件 consul dtm
 ## 启动consul     
    consul agent -dev
 ## 启动dtm        dtm -c dtm.yml      
    dtm.yml的内容如下：
    MicroService:
        Driver: 'dtm-driver-gozero' # 配置dtm使用go-zero的微服务协议
        Target: 'consul://localhost:8500/dtmservice' # 把dtm注册到etcd的这个地址
        EndPoint: 'localhost:36790' # dtm的本地地址

# xa模式 
    userrpc 要注册到consul      _ = consul.RegisterService(c.ListenOn, c.Consul)
    orderrpc要注册到consul       _ = consul.RegisterService(c.ListenOn, c.Consul)
