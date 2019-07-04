package service

import (
	"context"
)

type ProductService interface {
	GetProducts(ctx context.Context) (string, error)
}

type productService struct{}

func NewProdutService() ProductService {
	return productService{}
}

// Get will return today's date
func (productService) GetProducts(ctx context.Context) (string, error) {
	return "Products Successfully Loaded", nil
}
