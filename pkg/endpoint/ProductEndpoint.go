package endpoint

import (
	"context"
	"errors"

	transport "qasir-supplier/inventory/pkg/request/http" // part of transport
	"qasir-supplier/inventory/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

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
