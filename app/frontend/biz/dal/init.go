package dal

import (
	"github.coom/xmhu2001/gomall/app/frontend/biz/dal/mysql"
	"github.coom/xmhu2001/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
