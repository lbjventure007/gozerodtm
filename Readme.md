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

# saga 模式 本机两个库 模拟多库表操作 注意得使用子屏障 不然自己得处理空回滚 悬挂 幂等问题
    用户余额转入转出 两个库的用户表 user.user  user1.user
    sql如下： 
    CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
    `phone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `question` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '找回密码问题',
    `answer` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '找回密码答案',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `balance` decimal(16,2) DEFAULT NULL,
    `forzen_balance` decimal(16,2) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
    ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
     
    请求：
    Post: localhost:8884/user/tran-inout
    Body: {
            "inid":1,
            "outid":2,
            "balance":1
         }