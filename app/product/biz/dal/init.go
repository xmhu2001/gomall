package dal

import (
	"github.com/xmhu2001/gomall/app/product/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
