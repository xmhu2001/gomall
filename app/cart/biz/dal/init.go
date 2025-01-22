package dal

import (
	"github.com/xmhu2001/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
