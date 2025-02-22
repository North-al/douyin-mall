package service

import (
	"context"
	"fmt"
	"time"

	"github.com/North-al/douyin-mall/app/auth/biz/dal/redis"
	"github.com/North-al/douyin-mall/app/auth/conf"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	claims := jwt.MapClaims{
		"user_id": req.UserId,
		"exp":     time.Now().Add(time.Hour * time.Duration(conf.GetConf().JWT.ExpireTime)).Unix(), // 过期时间
		"iat":     time.Now().Unix(),                                                               // 创建时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.GetConf().JWT.Secret))
	if err != nil {
		return nil, err
	}

	redis.RedisClient.Set(s.ctx, fmt.Sprintf("user_token_%d", req.UserId), tokenString, time.Hour*time.Duration(conf.GetConf().JWT.ExpireTime))

	return &auth.DeliveryResp{Token: tokenString}, nil
}
