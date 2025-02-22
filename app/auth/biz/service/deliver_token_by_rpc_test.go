package service

import (
	"context"
	"testing"

	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

var deliveredToken string

func TestDeliverTokenByRPC_Run(t *testing.T) {
	InitTestEnv()

	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("deliver token failed: %v", err)
	}

	deliveredToken = resp.Token
	t.Logf("delivered token: %v", deliveredToken)
}
