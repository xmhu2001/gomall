package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/app/frontend/conf"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order/orderservice"
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
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
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
