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

func TestUserInfo_Run(t *testing.T) {

	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	dal.Init()

	ctx := context.Background()
	userQuery := query.NewUserQuery(ctx, mysql.DB)

	ctx = context.WithValue(ctx, "userId", int32(7))
	s := NewUserInfoService(ctx, userQuery)
	// init req and assert value

	req := &user.UserInfoReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3MTE2ODIsImlhdCI6MTczOTYyNTI4MiwidXNlcl9pZCI6N30.rfecUUKc4kw8tDhGdtNex5sKMtFOyzTszrE5jzzUxsg",
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
