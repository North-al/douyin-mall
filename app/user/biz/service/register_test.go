package service

import (
	"context"
	"os"
	"testing"

	"github.com/North-al/douyin-mall/app/user/biz/dal"
	"github.com/North-al/douyin-mall/app/user/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/user/biz/query"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {

	// 设置测试环境
	err := os.Setenv("GO_ENV", "test")
	if err != nil {
		t.Fatalf("设置环境变量失败: %v", err)
	}

	// 切换工作目录
	err = os.Chdir("../../")
	if err != nil {
		t.Fatalf("切换工作目录失败: %v", err)
	}

	// 初始化数据库连接
	dal.Init()

	// 验证数据库连接是否成功
	if mysql.DB == nil {
		t.Fatal("数据库连接为空")
	}

	ctx := context.Background()
	userQuery := query.NewUserQuery(mysql.DB)
	s := NewRegisterService(ctx, userQuery)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test7@test.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}

	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("注册失败: %v", err)
	}
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
