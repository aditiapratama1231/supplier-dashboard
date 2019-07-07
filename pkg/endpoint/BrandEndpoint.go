package endpoint

import (
	"context"
	"errors"

	payload "qasir-supplier/inventory/pkg/request/payload" // part of transport
	"qasir-supplier/inventory/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetProductsEndpoint returns the response from our service "get"
func MakeGetBrandsEndpoint(srv service.BrandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(payload.GetBrandsRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetBrands(ctx)
		if err != nil {
			return payload.GetBrandsResponse{d, err.Error()}, nil
		}
		return payload.GetBrandsResponse{d, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) GetBrands(ctx context.Context) (string, error) {
	req := payload.GetBrandsRequest{}
	resp, err := e.GetBrandsEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(payload.GetBrandsResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Brands, nil
}
