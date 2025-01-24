package service

import (
	"context"
	"testing"

	"github.com/North-al/douyin-mall/pkg"
	auth "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenSuccessByRPC_Run(t *testing.T) {
	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()

	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc3ODY3OTAsImlhdCI6MTczNzcwMDM5MCwidXNlcl9pZCI6MX0.biBPlCZOx7weDeN8dvfnHl5_eJSRUlOvZ644fBsvcyY",
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
