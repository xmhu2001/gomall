package dal

import (
	"github.com/xmhu2001/gomall/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
