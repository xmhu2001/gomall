package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/app/checkout/conf"
	"github.com/xmhu2001/gomall/app/checkout/utils"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
	})
}

func initProductClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Registry.RegistryAddress[0]})
	utils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	utils.MustHandleError(err)
}

func initCartClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Registry.RegistryAddress[0]})
	utils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	utils.MustHandleError(err)
}

func initPaymentClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Registry.RegistryAddress[0]})
	utils.MustHandleError(err)
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	utils.MustHandleError(err)
}
