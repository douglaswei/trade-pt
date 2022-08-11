package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/douglaswei/trade-pt/user/internal/config"
    "github.com/douglaswei/trade-pt/user/internal/model"
    "github.com/douglaswei/trade-pt/user/internal/query"
    "github.com/douglaswei/trade-pt/user/internal/server"
    "github.com/douglaswei/trade-pt/user/internal/svc"
    "github.com/douglaswei/trade-pt/user/user"
    "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

    "common/db"
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/core/service"
    "github.com/zeromicro/go-zero/zrpc"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)
    ctx := svc.NewServiceContext(c)
    svr := server.NewUserServer(ctx)
    //

    DB := db.LoadByDSN(c.MysqlConfig.DSN)
    query.Use(DB).User.WithContext(context.Background()).Create(&model.User{UserID: 10002, UserName: "Douglas", NikeName: "Dou"})
    s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
        user.RegisterUserServer(grpcServer, svr)

        if c.Mode == service.DevMode || c.Mode == service.TestMode {
            reflection.Register(grpcServer)
        }
    })

    _ = consul.RegisterService(c.ListenOn, c.Consul)

    defer s.Stop()

    fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
    s.Start()
}
