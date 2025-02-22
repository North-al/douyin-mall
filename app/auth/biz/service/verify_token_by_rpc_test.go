package service

import (
	"context"
	"testing"

	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenSuccessByRPC_Run(t *testing.T) {
	InitTestEnv()

	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{
		Token: deliveredToken, // 使用 deliver_test 中派发的 token
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
