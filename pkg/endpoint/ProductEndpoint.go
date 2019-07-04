package endpoint

import (
	"context"
	"errors"

	"qasir-supplier/inventory/pkg/service"
	transport "qasir-supplier/inventory/pkg/transport/http"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetEndpoint returns the response from our service "get"
func MakeGetProductsEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.GetProductsRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetProducts(ctx)
		if err != nil {
			return transport.GetProductsResponse{d, err.Error()}, nil
		}
		return transport.GetProductsResponse{d, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) GetProducts(ctx context.Context) (string, error) {
	req := transport.GetProductsRequest{}
	resp, err := e.GetProductsEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transport.GetProductsResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Products, nil
}
