package dal

import (
	"github.com/xmhu2001/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/xmhu2001/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
