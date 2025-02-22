package dal

import (
	"github.com/North-al/douyin-mall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
