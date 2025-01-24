package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/xmhu2001/gomall/app/payment/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/payment/biz/model"
	payment "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:   strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month: strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	if err = card.Validate(true); err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001,err.Error())
	}

	transactionId, err := uuid.NewRandom()

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001,err.Error())
	}
	// 记录支付记录
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId: 	 req.OrderId,
		TransactionId: transactionId.String(),
		Amount: 	  req.Amount,
		PayAt: time.Now(),
	})
	
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
