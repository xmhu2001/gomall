package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xmhu2001/gomall/app/cart/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/cart/biz/model"
	"github.com/xmhu2001/gomall/app/cart/infra/rpc"
	cart "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// check parameters
	produtcResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if produtcResp == nil || produtcResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	// add item to cart
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
