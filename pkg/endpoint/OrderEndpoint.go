package endpoint

import (
	"context"
	"fmt"

	"qasir-supplier/order/pkg/service"

	"github.com/go-kit/kit/endpoint"

	payload "qasir-supplier/order/pkg/request/payload"
)

func MakeGetMerchantOrdersEndpoint(srv service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.GetMerchantOrdersRequest)
		d := srv.GetMerchantOrders(req)
		return d, nil
	}
}

func MakeShowOrderDetailEndpoint(srv service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.GetItemsOrderRequest)
		d := srv.ShowOrderDetail(req)
		return d, nil
	}
}

func MakeCreateOrderEndpoint(srv service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateOrderRequest)
		d := srv.CreateOrder(req)
		return d, nil
	}
}

func MakeChangeOrderStatusEndpoint(srv service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ChangeOrderStatusRequest)
		d := srv.ChangeOrderStatus(req)
		return d, nil
	}
}

func (e Endpoints) GetMerchantOrders(ctx context.Context) (payload.GetMerchantOrdersResponse, error) {
	req := payload.GetMerchantOrdersRequest{}
	fmt.Println("Wewew")
	resp, _ := e.GetMerchantOrdersEndpoint(ctx, req)
	getResp := resp.(payload.GetMerchantOrdersResponse)
	return getResp, nil
}

func (e Endpoints) ShowOrderDetail(ctx context.Context) (payload.GetMerchantOrdersResponse, error) {
	req := payload.GetMerchantOrdersRequest{}
	resp, _ := e.ShowOrderDetailEndpoint(ctx, req)
	getResp := resp.(payload.GetMerchantOrdersResponse)
	return getResp, nil
}
