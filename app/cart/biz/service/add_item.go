package service

import (
	"context"

	"github.com/North-al/douyin-mall/app/cart/biz/model"
	"github.com/North-al/douyin-mall/app/cart/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/cart/rpc"
	cart "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/cart"
	product "github.com/North-al/douyin-mall/rpc_gen/kitex_gen/product"

	kerrors "github.com/cloudwego/kitex/pkg/kerrors"
	

)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	//验证产品的id是否真实存在,会调用product的代码
	productResp, err:=rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	//如果产品不存在,返回错误
	if productResp ==nil || productResp.Product.Id ==0 {
		return nil, kerrors.NewBizStatusError(40004,"product not found")
		
	}
	CartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}

	err = model.AddItem(s.ctx, mysql.DB, CartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
