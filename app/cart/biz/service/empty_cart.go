package service

import (
	"context"

	cart "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/cart"
	"github.com/North-al/douyin-mall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyCart( mysql.DB,s.ctx,req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001,err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}

