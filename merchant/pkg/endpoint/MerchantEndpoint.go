package endpoint

import (
	"context"
	"errors"

	model "qasir-supplier/merchant/pkg/request/model" // part of transport
	"qasir-supplier/merchant/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetProductsEndpoint returns the response from our service "get"
func MakeGetMerchantsEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(model.GetMerchantsRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetMerchants(ctx)
		if err != nil {
			return model.GetMerchantsResponse{d, err.Error()}, nil
		}
		return model.GetMerchantsResponse{d, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) GetMerchants(ctx context.Context) (string, error) {
	req := model.GetMerchantsRequest{}
	resp, err := e.GetMerchantsEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(model.GetMerchantsResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Merchants, nil
}
