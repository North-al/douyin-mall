package service

import (
	"context"

	"github.com/North-al/douyin-mall/app/cart/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/cart/biz/model"
	cart "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"

)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartByUserID(s.ctx,mysql.DB,req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002,err.Error())

	}
	// var items []*cart.CartItem
	// for _, item := range list {
	// 	items = append(items, &cart.CartItem{
	// 		ProductId: item.ProductId,
	// 		Quantity:  int32(item.Qty),
	// 	})
	// }
	// return &cart.GetCartResp{Items: items}, nil

	items := make([]*cart.CartItem, len(list))
	for i, v := range list {
		items[i] = &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  int32(v.Qty),
		}
	}
	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.GetUserId(),
			Items:  items,
		},
	}, nil

	//
}
