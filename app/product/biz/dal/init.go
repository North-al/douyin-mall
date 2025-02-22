package dal

import (
	"github.com/North-al/douyin-mall/app/product/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
