package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// VariantWholesaler model
type VariantWholesaler struct {
	gorm.Model
	ID              int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	VariantID       int64     `gorm:"column:variant_id" json:"variant_id"`
	QuantityMinimum float64   `gorm:"column:quantity_minimum" json:"quantity_minimum"`
	WholesalePrice  float64   `gorm:"column:wholesale_price" json:"wholesale_price"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
