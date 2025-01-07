package dal

import (
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
