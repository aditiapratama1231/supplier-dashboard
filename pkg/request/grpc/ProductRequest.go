package request

import (
	"context"
	"qasir-supplier/inventory/pb"

	transport "qasir-supplier/inventory/pkg/request/http"
)

func EncodeGRPCProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	return &pb.ProductRequest{}, nil
}

func DecodeGRPCProductRequest(ctx context.Context, r interface{}) (interface{}, error) {
	return transport.GetProductsRequest{}, nil
}

func EncodeGRPCProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(transport.GetProductsResponse)
	return &pb.ProductResponse{
		Message: resp.Products,
		Err:     resp.Err,
	}, nil
}

func DecodeGRPCProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.ProductResponse)
	return transport.GetProductsResponse{
		Products: resp.Message,
		Err:      resp.Err,
	}, nil
}
