package dal

import (
	"github.com/North-al/douyin-mall/app/user/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
