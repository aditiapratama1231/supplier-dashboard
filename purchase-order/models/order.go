package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order model
type Order struct {
	gorm.Model
	VariantID      int64       `gorm:"column:variant_id" json:"variant_id,omitempty"`
	SupplierID     int64       `gorm:"column:supplier_id" json:"supplier_id,omitempty"`
	MerchantID     int64       `gorm:"column:merchant_id" json:"merchant_id,omitempty"`
	OutletID       int64       `gorm:"column:outlet_id" json:"outlet_id,omitempty"`
	Subtotal       float64     `gorm:"column:subtotal" json:"subtotal,omitempty"`
	DiscountAmount float64     `gorm:"column:discount_amount" json:"discount_amount,omitempty"`
	GrandTotal     float64     `gorm:"column:grand_total" json:"grand_total,omitempty"`
	OrderNo        string      `gorm:"column:order_no" json:"order_no,omitempty"`
	Status         string      `gorm:"column:status" json:"status,omitempty"`
	HasInvoice     bool        `gorm:"column:has_invoice" json:"has_invoice,omitempty"`
	CanceledAt     time.Time   `json:"canceled_at" db:"canceled_at,omitempty" sql:"DEFAULT:'current_timestamp'"`
	ShippedAt      time.Time   `json:"shipped_at" db:"shipped_at,omitempty" sql:"DEFAULT:'current_timestamp'"`
	DeliveredAt    time.Time   `json:"delivered_at" db:"delivered_at,omitempty" sql:"DEFAULT:'current_timestamp'"`
	CompletedAt    time.Time   `json:"completed_at" db:"completed_at,omitempty" sql:"DEFAULT:'current_timestamp'"`
	OrderItems     []OrderItem `json:"order_items,omitempty,omitempty" sql:"DEFAULT:'current_timestamp'"`
}
