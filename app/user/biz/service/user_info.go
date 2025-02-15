package service

import (
	"context"

	"github.com/North-al/douyin-mall/app/user/biz/query"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
)

type UserInfoService struct {
	ctx       context.Context
	userQuery *query.UserQuery
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context, userQuery *query.UserQuery) *UserInfoService {
	return &UserInfoService{ctx: ctx, userQuery: userQuery}
}

// Run create note info
func (s *UserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	return
}
