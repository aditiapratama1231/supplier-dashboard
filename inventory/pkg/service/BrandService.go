package service

import (
	"context"

	"github.com/jinzhu/gorm"
)

type BrandService interface {
	GetBrands(ctx context.Context) (string, error)
}

type brandService struct{}

func NewBrandService(db *gorm.DB) BrandService {
	query = db
	return brandService{}
}

// Get will return today's date
func (brandService) GetBrands(ctx context.Context) (string, error) {
	return "Brands Successfully Loaded", nil
}
