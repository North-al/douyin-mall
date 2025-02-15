package service

import (
	"context"
	"testing"

	"github.com/North-al/douyin-mall/app/auth/biz/dal"
	"github.com/North-al/douyin-mall/pkg"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenSuccessByRPC_Run(t *testing.T) {
	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	// dal.Init()

	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3MDIzODksImlhdCI6MTczOTYxNTk4OSwidXNlcl9pZCI6MX0.7d0xii0q_Uo4yo1zeNKjTkaSUfbsMbDwY2aPPU54v88",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestVerifyTokenErrorByRPC_Run(t *testing.T) {
	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	dal.Init()

	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUz",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
