package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model         //using this will generated id, (timestamp) automatically
	CategoryID  int64  `gorm:"column:category_id" json:"category_id"`
	BrandID     int64  `gorm:"column:brand_id" json:"brand_id"`
	MerchantID  int64  `gorm:"column:merchant_id" json:"merchant_id"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Image       string `gorm:"column:image" json:"image"`
	HasText     int8   `gorm:"column:has_text" json:"has_text"`
	CreatedBy   *int64 `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   *int64 `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy   *int64 `gorm:"column:deleted_by" json:"deleted_by"`
	IsPublished int8   `gorm:"column:is_published" json:"is_published"`
}

// Insert Product To Database
func InsertProduct(db *gorm.DB, p *Product) (err error) {
	if err = db.Save(p).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
