package request

// import (
// 	"context"
// 	"qasir-supplier/inventory/pb"

// 	model "qasir-supplier/inventory/pkg/request/model"
// )

// func EncodeGRPCProductRequest(_ context.Context, r interface{}) (interface{}, error) {
// 	return &pb.ProductRequest{}, nil
// }

// func DecodeGRPCProductRequest(ctx context.Context, r interface{}) (interface{}, error) {
// 	return model.GetProductsRequest{}, nil
// }

// func EncodeGRPCProductResponse(_ context.Context, r interface{}) (interface{}, error) {
// 	resp := r.(model.GetProductsResponse)
// 	return &pb.ProductResponse{
// 		Products: resp.Products,
// 		Err:      resp.Err,
// 	}, nil
// }

// func DecodeGRPCProductResponse(_ context.Context, r interface{}) (interface{}, error) {
// 	resp := r.(*pb.ProductResponse)
// 	return model.GetProductsResponse{
// 		Products: resp.Message,
// 		Err:      resp.Err,
// 	}, nil
// }
