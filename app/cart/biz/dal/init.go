package dal

import (
	"github.com/North-al/douyin-mall/app/cart/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
