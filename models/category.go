package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Category model
type Category struct {
	gorm.Model
	ID        int64     `gorm:"column:id" gorm:"primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Slug      string    `gorm:"column:slug" json:"slug"`
	CreatedBy int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy int64     `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy int64     `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Parent    int64     `gorm:"column:parent" json:"parent"`
}
