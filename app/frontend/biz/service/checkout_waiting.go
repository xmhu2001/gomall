package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
	rpccheckout "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardCvv:             req.Cvv,
		},
	})

	if err != nil {
		return nil, err
	}

	// 助手函数 WarpResp 用于包装响应
	return utils.H{
		"title":    "Checkout_waiting",
		"redirect": "/checkout/result",
	}, nil
}
