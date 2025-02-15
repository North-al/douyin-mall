package service

import (
	"context"
	"testing"

	"github.com/North-al/douyin-mall/app/user/biz/dal"
	"github.com/North-al/douyin-mall/app/user/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/user/biz/query"
	"github.com/North-al/douyin-mall/pkg"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {

	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	dal.Init()

	ctx := context.Background()
	userQuery := query.NewUserQuery(ctx, mysql.DB)
	s := NewRegisterService(ctx, userQuery)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test8@test.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}

	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("注册失败: %v", err)
	}
	t.Logf("resp: %v", resp)
}
