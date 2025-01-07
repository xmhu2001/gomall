package dal

import (
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
