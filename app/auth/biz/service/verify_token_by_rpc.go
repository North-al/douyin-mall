package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/North-al/douyin-mall/app/auth/biz/dal/redis"
	"github.com/North-al/douyin-mall/app/auth/conf"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp = &auth.VerifyResp{}
	token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(conf.GetConf().JWT.Secret), nil
	})

	if err != nil {
		resp.Res = false
		return resp, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	result, err := redis.RedisClient.Get(s.ctx, fmt.Sprintf("user_token_%d", (int)(claims["user_id"].(float64)))).Result()
	if err != nil {
		return nil, errors.New(fmt.Errorf("get token from redis failed, err: %v", err).Error())
	}

	fmt.Println("result: ", result)
	fmt.Println("req.Token: ", req.Token)

	if result != req.Token {
		resp.Res = false
		return resp, errors.New("token is invalid")
	}

	if ok && token.Valid {
		resp.Res = true
		return resp, nil
	}

	resp.Res = false
	return resp, err
}
