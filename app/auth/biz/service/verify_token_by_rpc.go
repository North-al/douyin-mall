package service

import (
	"context"
	"fmt"

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
	// TODO: 缺少redis对比token信息

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

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		resp.Res = true
		return resp, nil
	}

	resp.Res = false
	return resp, err
}
