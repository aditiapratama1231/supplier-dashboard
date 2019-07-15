package models

import (
	"github.com/jinzhu/gorm"
)

// OrderStatusHistory model
type OrderStatusHistory struct {
	gorm.Model
	MerchantID      int64  `gorm:"column:merchant_id" json:"merchant_id,omitempty"`
	OrderID         int64  `gorm:"column:order_id" json:"order_id,omitempty"`
	OrderStatusCode string `gorm:"column:order_status_code" json:"order_status_code,omitempty"`
	StatusName      string `gorm:"column:status_name" json:"status_name,omitempty"`
}
