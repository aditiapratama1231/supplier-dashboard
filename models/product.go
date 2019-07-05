package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model
	ID          int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	CategoryID  int64     `gorm:"column:category_id" json:"category_id"`
	BrandID     int64     `gorm:"column:brand_id" json:"brand_id"`
	MerchantID  int64     `gorm:"column:merchant_id" json:"merchant_id"`
	ProductCode string    `gorm:"column:product_code" json:"product_code"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Image       string    `gorm:"column:image" json:"image"`
	HasTax      int8      `gorm:"column:has_tax" json:"has_tax"`
	CreatedBy   int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   int64     `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy   int64     `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	IsPublished int8      `gorm:"column:is_published" json:"is_published"`
}
