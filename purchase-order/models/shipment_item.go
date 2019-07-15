package models

import (
	"github.com/jinzhu/gorm"
)

// ShipmentItem model
type ShipmentItem struct {
	gorm.Model
	ShipmentID  int64   `gorm:"column:shipment_id" json:"shipment_id,omitempty"`
	OrderItemID int64   `gorm:"column:order_item_id" json:"order_item_id,omitempty"`
	VariantID   int64   `gorm:"column:variant_id" json:"variant_id,omitempty"`
	Price       float64 `gorm:"column:price" json:"price,omitempty"`
	Qty         float64 `gorm:"column:qty" json:"qty,omitempty"`
	Weight      float64 `gorm:"column:weight" json:"weight,omitempty"`
	Name        string  `gorm:"column:name" json:"name,omitempty"`
	Sku         string  `gorm:"column:sku" json:"sku,omitempty"`
	Volume      float64 `gorm:"column:volume" json:"volume,omitempty"`
}
