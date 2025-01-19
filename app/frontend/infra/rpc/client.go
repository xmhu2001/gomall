package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xmhu2001/gomall/app/frontend/conf"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Hertz.RegistryAddr})
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
