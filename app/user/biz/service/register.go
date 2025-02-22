package service

import (
	"context"
	"errors"

	"github.com/North-al/douyin-mall/app/user/biz/query"
	user "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx       context.Context
	userQuery *query.UserQuery
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context, userQuery *query.UserQuery) *RegisterService {
	return &RegisterService{ctx: ctx, userQuery: userQuery}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 1. 检查密码是否符合规范
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}

	// 2. 检查用户是否存在
	_, err = s.userQuery.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser, err := s.userQuery.CreateUser(req.Email, req.Password)
			if err != nil {
				return nil, err
			}
			resp = &user.RegisterResp{
				UserId: int32(newUser.ID),
			}

			return resp, nil
		}

		return nil, err
	}

	return nil, errors.New("user already exists")
}
