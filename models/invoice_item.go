package models

import (
	"github.com/jinzhu/gorm"
)

// InvoiceItem model
type InvoiceItem struct {
	gorm.Model
	InvoiceID       int64   `gorm:"column:invoice_id" json:"invoice_id,omitempty"`
	VariantID       int64   `gorm:"column:variant_id" json:"variant_id,omitempty"`
	Qty             float64 `gorm:"column:qty" json:"qty,omitempty"`
	Price           float64 `gorm:"column:price" json:"price,omitempty"`
	Sku             string  `gorm:"column:sku" json:"sku,omitempty"`
	Name            string  `gorm:"column:name" json:"name,omitempty"`
	InvoiceItemscol string  `gorm:"column:privileges" json:"privileges,omitempty"`
}

// // Insert Product To Database
// func InsertProduct(db *gorm.DB, p *Product) (err error) {
// 	if err = db.Save(p).Error; err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil
// }
