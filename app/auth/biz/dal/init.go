package dal

import (
	"github.com/North-al/douyin-mall/app/auth/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
