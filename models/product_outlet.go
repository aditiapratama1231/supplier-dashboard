package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductOutlet struct {
	gorm.Model
	ID         int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	ProductsID int64     `gorm:"column:products_id" json:"products_id"`
	OutletsID  int64     `gorm:"column:outlets_id" json:"outlets_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
