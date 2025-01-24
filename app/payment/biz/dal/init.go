package dal

import (
	"github.com/xmhu2001/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
