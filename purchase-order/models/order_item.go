package models

import (
	"github.com/jinzhu/gorm"
)

// OrderItem model
type OrderItem struct {
	gorm.Model
	VariantID     int64   `gorm:"column:variant_id" json:"variant_id,omitempty"`
	OrderID       int64   `gorm:"column:order_id" json:"order_id,omitempty"`
	Qty           float64 `gorm:"column:qty" json:"qty,omitempty"`
	QtyInvoice    float64 `gorm:"column:qty_invoice" json:"qty_invoice,omitempty"`
	QtyShipped    float64 `gorm:"column:qty_shipped" json:"qty_shipped,omitempty"`
	QtyCanceled   float64 `gorm:"column:qty_canceled" json:"qty_canceled,omitempty"`
	QtyRefunded   float64 `gorm:"column:qty_refunded" json:"qty_refunded,omitempty"`
	PriceSell     float64 `gorm:"column:price_sell" json:"price_sell,omitempty"`
	PriceSellUnit float64 `gorm:"column:price_sell_unit" json:"price_sellUnit,omitempty"`
	PriceBase     float64 `gorm:"column:price_base" json:"price_base,omitempty"`
	PriceBaseUnit float64 `gorm:"column:price_base_unit" json:"price_base_unit,omitempty"`
	Sku           string  `gorm:"column:sku" json:"sku,omitempty"`
	Name          string  `gorm:"column:name" json:"name,omitempty"`
}
