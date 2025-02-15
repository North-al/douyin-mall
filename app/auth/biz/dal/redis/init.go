package redis

import (
	"context"
	"fmt"

	"github.com/North-al/douyin-mall/app/auth/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})

	klog.Info("RedisClient: ", RedisClient)

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("Init redis failed: %v", err))
	}
}
