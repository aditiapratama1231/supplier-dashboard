package models

import (
	"github.com/jinzhu/gorm"
)

// Invoice model
type Invoice struct {
	gorm.Model
	OrderID       int64   `gorm:"column:order_id" json:"order_id,omitempty"`
	MerchantID    int64   `gorm:"column:merchant_id" json:"merchant_id,omitempty"`
	OutletID      int64   `gorm:"column:OutletID" json:"OutletID,omitempty"`
	InvoiceNo     string  `gorm:"column:invoice_no" json:"invoice_no,omitempty"`
	Status        int8    `gorm:"column:sku" json:"sku,omitempty"`
	TotalQty      float64 `gorm:"column:total_qty" json:"total_qty,omitempty"`
	TotalAmount   float64 `gorm:"column:total_amount" json:"total_amount,omitempty"`
	TotalDiscount float64 `gorm:"column:total_discount" json:"total_discount,omitempty"`
}
