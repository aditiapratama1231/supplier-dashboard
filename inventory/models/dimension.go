package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Dimension struct {
	gorm.Model
	ID        int64      `gorm:"column:id" gorm:"primary_key" json:"id"`
	VariantID int64      `gorm:"column:variant_id" json:"variant_id"`
	Width     float64    `gorm:"column:width" json:"width"`
	Long      float64    `gorm:"column:long" json:"long"`
	Height    float64    `gorm:"column:height" json:"height"`
	Weight    float64    `gorm:"column:height" json:"height"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}
