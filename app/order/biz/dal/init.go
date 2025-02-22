package dal

import (
	"github.com/North-al/douyin-mall/app/order/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
