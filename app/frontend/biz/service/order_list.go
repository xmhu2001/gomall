package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	"github.com/xmhu2001/gomall/app/frontend/types"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	rpcorder "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {

	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	// Get order list by rpc call
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: uint32(userId)})

	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, oi := range v.Items {
			pr, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: oi.Item.ProductId})
			if err != nil {
				return nil, err
			}
			if pr == nil || pr.Product == nil {
				continue
			}
			items = append(items, types.OrderItem{
				ProductName: pr.Product.Name,
				Picture:     pr.Product.Picture,
				Qty:         oi.Item.Quantity,
				Cost:        oi.Cost,
			})
			total += oi.Cost
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:    v.OrderId,
			CreateDate: created.Format("2006-01-02 15:04:05"),
			Cost:       total,
			Items:      items,
		})
	}

	return utils.H{"title": "Order List", "orders": list}, nil
}
