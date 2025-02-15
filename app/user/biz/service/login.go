package service

import (
	"context"
	"errors"

	"github.com/North-al/douyin-mall/app/user/biz/query"
	"github.com/North-al/douyin-mall/app/user/biz/rpc"
	"github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
	"github.com/kr/pretty"
)

type LoginService struct {
	ctx       context.Context
	userQuery *query.UserQuery
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context, userQuery *query.UserQuery) *LoginService {
	return &LoginService{ctx: ctx, userQuery: userQuery}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {

	resp = &user.LoginResp{}
	queryUser, err := s.userQuery.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if queryUser.Password != req.Password {
		return nil, errors.New("账号或密码错误")
	}

	// 调用auth客户端生成token
	authResp, err := rpc.GetAuthClient().DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(queryUser.ID),
	})
	if err != nil {
		return nil, err
	}

	pretty.Printf("authResp: %+v\n", authResp)

	resp.UserId = int32(queryUser.ID)

	return
}
