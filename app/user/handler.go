package main

import (
	"context"

	"github.com/North-al/douyin-mall/app/auth/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/user/biz/query"
	"github.com/North-al/douyin-mall/app/user/biz/service"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx, query.NewUserQuery(ctx, mysql.DB)).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx, query.NewUserQuery(ctx, mysql.DB)).Run(req)

	return resp, err
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	resp, err = service.NewUserInfoService(ctx).Run(req)

	return resp, err
}
