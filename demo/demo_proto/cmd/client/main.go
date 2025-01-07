package main

import (
	"context"
	"github.com/cloudwego/kitex/transport"
	"github.com/xmhu2001/gomall/demo/demo_proto/kitex_gen/pbapi"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"

	"github.com/kitex-contrib/registry-etcd"
	echo "github.com/xmhu2001/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
)

func main() {
	// 创建 etcd 注册中心解析器
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	// 创建客户端选项
	opts := []client.Option{
		client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithRPCTimeout(3 * time.Second),
	}

	// 创建客户端实例
	cl := echo.MustNewClient("demo_proto", opts...)

	// 调用服务方法
	for {
		resp, err := cl.Echo(context.TODO(), &pbapi.Request{Message: "hello world"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Message)
		time.Sleep(3 * time.Second)
	}
}
