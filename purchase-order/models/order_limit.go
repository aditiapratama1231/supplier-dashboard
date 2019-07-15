package models

import (
	"github.com/jinzhu/gorm"
)

// OrderLimit model
type OrderLimit struct {
	gorm.Model
	VariantID int64   `gorm:"column:variant_id" json:"variant_id,omitempty"`
	MinQty    float64 `gorm:"column:min_qty" json:"min_qty,omitempty"`
	MaxQty    float64 `gorm:"column:max_qty" json:"max_qty,omitempty"`
}
