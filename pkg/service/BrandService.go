package service

import (
	"context"
)

type BrandService interface {
	GetBrands(ctx context.Context) (string, error)
}

type brandService struct{}

func NewBrandService() BrandService {
	return brandService{}
}

// Get will return today's date
func (brandService) GetBrands(ctx context.Context) (string, error) {
	return "Brands Successfully Loaded", nil
}
