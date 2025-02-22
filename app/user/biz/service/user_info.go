package service

import (
	"context"
	"fmt"

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

	userId, ok := s.ctx.Value("userId").(int32)
	if !ok {
		return nil, fmt.Errorf("failed to get userId from context")
	}

	queryUser, err := s.userQuery.GetUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}

	resp = &user.UserInfoResp{
		Id:        int32((queryUser.ID)),
		Username:  queryUser.Username,
		Email:     queryUser.Email,
		Avatar:    queryUser.Avatar,
		CreatedAt: queryUser.CreatedAt.String(),
		UpdatedAt: queryUser.UpdatedAt.String(),
	}

	return resp, nil
}
