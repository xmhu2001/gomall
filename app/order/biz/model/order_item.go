package model

import "gorm.io/gorm"

// 订单商品关联表
type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11)"`
	OrderIdRefer string  `gorm:"type:varchar(100);index"`
	Quantity     uint32  `gorm:"type:int(11)"`
	Cost         float32 `gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
