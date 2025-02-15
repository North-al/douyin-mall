package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/North-al/douyin-mall/app/user/biz/rpc"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

// VerifyToken 验证 token 的公共函数
func VerifyToken(ctx context.Context, token string) (int32, error) {
	verifyResp, err := rpc.GetAuthClient().VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return 0, fmt.Errorf("token verification failed: %v", err)
	}

	if !verifyResp.Res {
		return 0, errors.New("invalid token")
	}

	// 从 token 中获取 user ID
	// userId := verifyResp.Res
	return 1, nil
}
