package endpoint

import (
	"context"
	"errors"

	model "qasir-supplier/inventory/pkg/request/model" // part of transport
	"qasir-supplier/inventory/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetProductsEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(model.GetProductsRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetProducts(ctx)
		if err != nil {
			return model.GetProductsResponse{d, err.Error()}, nil
		}
		return model.GetProductsResponse{d, ""}, nil
	}
}

func (e Endpoints) GetProducts(ctx context.Context) (string, error) {
	req := model.GetProductsRequest{}
	resp, err := e.GetProductsEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(model.GetProductsResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Products, nil
}
