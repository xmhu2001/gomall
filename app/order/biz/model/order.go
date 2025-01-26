package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

// 定义订单的表结构
type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`              // 订单号
	UserId       uint32      `gorm:"int(11)"`                                    // 用户ID
	UserCurrency string      `gorm:"type:varchar(10)"`                           // 用户货币
	Consignee    Consignee   `gorm:"embedded"`                                   // 收货人信息
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"` // 订单商品
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
