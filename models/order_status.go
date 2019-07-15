package models

import (
	"github.com/jinzhu/gorm"
)

//OrderStatus model
type OrderStatus struct {
	gorm.Model
	Code  string `gorm:"column:code" json:"code,omitempty"`
	Name  string `gorm:"column:name" json:"name,omitempty"`
	Order int8   `gorm:"column:order" json:"order,omitempty"`
}
