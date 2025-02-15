package service

import (
	"context"
	"testing"

	"github.com/North-al/douyin-mall/app/auth/biz/dal"
	"github.com/North-al/douyin-mall/pkg"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	dal.Init()

	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
