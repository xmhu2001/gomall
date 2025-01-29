package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/app/frontend/conf"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	// 更新具体规则
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r), client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(
		fallback.UnwrapHelper(
			func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
				if err == nil {
					return resp, nil
				}
				methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
				if methodName != "ListProducts" {
					return resp, err
				}
				return &product.ListProductsResp{
					Products: []*product.Product{
						{
							Price:       5.6,
							Id:          3,
							Picture:     "/static/image/manga.png",
							Name:        "manga-1",
							Description: "manga",
						},
					},
				}, nil
			}),
	)))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
