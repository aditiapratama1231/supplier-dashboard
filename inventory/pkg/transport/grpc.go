package transport

// import (
// 	"context"
// 	"qasir-supplier/inventory/pb"

// 	grpctransport "github.com/go-kit/kit/transport/grpc"

// 	"qasir-supplier/inventory/pkg/endpoint"

// 	transport "qasir-supplier/inventory/pkg/request/grpc"
// )

// type grpcServer struct {
// 	getProducts grpctransport.Handler
// }

// func (s *grpcServer) GetProducts(ctx context.Context, r *pb.ProductRequest) (*pb.ProductResponse, error) {
// 	_, resp, err := s.getProducts.ServeGRPC(ctx, r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.(*pb.ProductResponse), nil
// }

// func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.ProductsServer {
// 	return &grpcServer{
// 		getProducts: grpctransport.NewServer(
// 			endpoints.GetProductsEndpoint,
// 			transport.DecodeGRPCProductRequest,
// 			transport.EncodeGRPCProductResponse,
// 		),
// 	}
// }
