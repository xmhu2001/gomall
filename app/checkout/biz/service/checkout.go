package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xmhu2001/gomall/app/checkout/infra/rpc"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/order"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/payment"
	"github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "empty cart")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)
	for _, item := range cartResult.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		total += float32(item.Quantity) * productResp.Product.Price
		oi = append(oi, &order.OrderItem{
			Item: item,
			Cost: float32(item.Quantity) * productResp.Product.Price,
		})
	}

	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			Street:  req.Address.StreetAddress,
			City:    req.Address.City,
			State:   req.Address.State,
			Country: req.Address.Country,
			ZipCode: req.Address.ZipCode,
		},
		Items: oi,
	})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	var orderId string
	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})

	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return resp, nil
}
