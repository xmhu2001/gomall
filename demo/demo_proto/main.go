package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal"
	"github.com/xmhu2001/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/xmhu2001/gomall/demo/demo_proto/middleware"
	"log"
	"net"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	dal.Init()
	opts := kitexInit()

	svr := echo.NewServer(new(EchoImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithMiddleware(middleware.Middleware))

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))
	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo_proto",
	}))
	return
}
