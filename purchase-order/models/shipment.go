package models

import (
	"github.com/jinzhu/gorm"
)

// Shipment model
type Shipment struct {
	gorm.Model
	OrderID      int64   `gorm:"column:order_id" json:"order_id,omitempty"`
	OutletID     int64   `gorm:"column:outlet_id" json:"outlet_id,omitempty"`
	TotalWeight  float64 `gorm:"column:total_weight" json:"total_weight,omitempty"`
	TotalQty     float64 `gorm:"column:total_qty" json:"total_qty,omitempty"`
	ShippingNo   string  `gorm:"column:shipping_no" json:"shipping_no,omitempty"`
	ShippingNote string  `gorm:"column:shipping_note" json:"shipping_note,omitempty"`
	CustomerNote string  `gorm:"column:customer_note" json:"customer_note,omitempty"`
	Status       int8    `gorm:"column:status" json:"status,omitempty"`
}
