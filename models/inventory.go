package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Inventory model
type Inventory struct {
	gorm.Model
	ID         int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	TrackStock int8      `gorm:"column:track_stock" json:"track_stock"`
	AlertStock int8      `gorm:"column:alert_stock" json:"alert_stock"`
	StockMin   float64   `gorm:"column:stock_min" json:"stock_min"`
	StockIn    float64   `gorm:"column:stock_in" json:"stock_in"`
	StockOut   float64   `gorm:"column:stock_out" json:"stock_out"`
	CreatedBy  int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy  int64     `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy  int64     `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
