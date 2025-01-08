package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/xmhu2001/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/xmhu2001/gomall/demo/demo_proto/middleware"
	"log"

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
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	}

	// 创建客户端实例
	cl := echo.MustNewClient("demo_proto", opts...)

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	// 调用服务方法
	resp, err := cl.Echo(ctx, &pbapi.Request{Message: "hello world"})
	if err != nil {
		var bizErr *kerrors.GRPCBizStatusError
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Println(bizErr)
		}
		log.Fatal(err)
	}
	log.Println(resp.Message)
}
