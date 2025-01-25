package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	rpccart "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string

	// Get the user_id
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	// Get the cart items
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})

	if err != nil || cartResp == nil {
		return nil, err
	}

	// Get the product details and total price
	var total float32
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: item.ProductId})

		if err != nil {
			return nil, err
		}
		if productResp == nil {
			continue
		}

		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 32),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(item.Quantity)),
		})
		total += float32(item.Quantity) * p.Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
