package service

import (
	"context"
)

type MerchantService interface {
	GetMerchants(ctx context.Context) (string, error)
}

type merchantService struct{}

func NewMerchantService() MerchantService {
	return merchantService{}
}

// Get will return today's date
func (merchantService) GetMerchants(ctx context.Context) (string, error) {
	return "Merchants Successfully Loaded", nil
}
