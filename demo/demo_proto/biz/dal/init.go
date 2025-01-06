package dal

import (
	"github.coom/xmhu2001/gomall/demo/demo_proto/biz/dal/mysql"
	"github.coom/xmhu2001/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
