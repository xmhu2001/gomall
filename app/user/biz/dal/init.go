package dal

import (
	"github.com/xmhu2001/gomall/app/user/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
