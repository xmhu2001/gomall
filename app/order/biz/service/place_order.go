package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/xmhu2001/gomall/app/order/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/order/biz/model"
	order "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// check params
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "items empty")
		return nil, err
	}
	// 事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// generate order_id
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			o.Consignee.StreetAddress = req.Address.Street
			o.Consignee.City = req.Address.City
			o.Consignee.State = req.Address.State
			o.Consignee.Country = req.Address.Country
			o.Consignee.ZipCode = req.Address.ZipCode
		}
		// 写入
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		// 订单表写入成功，则再去写订单商品的子表
		var orderItems []model.OrderItem
		for _, v := range req.Items {
			orderItems = append(orderItems, model.OrderItem{
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				OrderIdRefer: orderId.String(),
				Cost:         v.Cost,
			})
		}

		if err = tx.Create(orderItems).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{OrderId: orderId.String()},
		}

		return nil
	})

	return
}
