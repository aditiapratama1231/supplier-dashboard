package endpoint

import (
	"context"
	"errors"

	"qasir-supplier/inventory/pkg/service"
	transport "qasir-supplier/inventory/pkg/transport/http"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetProductsEndpoint returns the response from our service "get"
func MakeGetBrandsEndpoint(srv service.BrandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.GetBrandsRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetBrands(ctx)
		if err != nil {
			return transport.GetBrandsResponse{d, err.Error()}, nil
		}
		return transport.GetBrandsResponse{d, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) GetBrands(ctx context.Context) (string, error) {
	req := transport.GetBrandsRequest{}
	resp, err := e.GetBrandsEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transport.GetBrandsResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Brands, nil
}
