package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// gorm 的公共 model
type PaymentLog struct {
	gorm.Model
	UserId uint32 `json:"user_id"`
	OrderId string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
	Amount float32 `json:"amount"`
	PayAt time.Time `json:"pay_at"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}