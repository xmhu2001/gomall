package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xmhu2001/gomall/app/cart/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/cart/biz/model"
	cart "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50003, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range list {
		items = append(items, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}
	return &cart.GetCartResp{Items: items}, nil
}
