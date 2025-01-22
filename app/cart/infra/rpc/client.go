package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/app/cart/conf"
	utils "github.com/xmhu2001/gomall/app/cart/utils"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Registry.RegistryAddress[0]})
	utils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	utils.MustHandleError(err)
}
