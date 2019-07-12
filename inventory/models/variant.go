package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Variant model
type Variant struct {
	gorm.Model
	ID             int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	ProductID      int64     `gorm:"column:product_id" json:"product_id"`
	MerchantID     int64     `gorm:"column:merchant_id" json:"merchant_id"`
	Name           string    `gorm:"column:name" json:"name"`
	Sku            string    `gorm:"column:sku" json:"sku"`
	PriceBase      float64   `gorm:"column:price_base" json:"price_base"`
	PriceSale      float64   `gorm:"column:price_sale" json:"price_sale"`
	BundleQuantity float64   `gorm:"column:bundle_quantity" json:"bundle_quantity"`
	VariantBundle  int64     `gorm:"column:variant_bundle" json:"variant_bundle"`
	VariantCode    string    `gorm:"column:variant_code" json:"variant_code"`
	IsIngredient   int8      `gorm:"column:is_ingredient" json:"is_ingredient"`
	IsRecipe       int8      `gorm:"column:is_recipe" json:"is_recipe"`
	IsBundle       int8      `gorm:"column:is_bundle" json:"is_bundle"`
	CreatedBy      int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy      int64     `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy      int64     `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
