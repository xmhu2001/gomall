package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xmhu2001/gomall/app/order/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/order/biz/model"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	orderlist, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var results []*order.Order
	for _, v := range orderlist {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		results = append(results, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Address: &order.Address{
				Street:  v.Consignee.StreetAddress,
				City:    v.Consignee.City,
				State:   v.Consignee.State,
				Country: v.Consignee.Country,
				ZipCode: v.Consignee.ZipCode,
			},
			Email: v.Consignee.Email,
			Items: items,
		})
	}

	resp = &order.ListOrderResp{
		Orders: results,
	}

	return resp, nil
}
