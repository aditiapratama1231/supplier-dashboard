package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	payload "qasir-supplier/order/pkg/request/payload"
)

func DecodeGetMerchantOrdersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetMerchantOrdersRequest
	return req, nil
}

func DecodeShowOrderDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// var req payload.GetItemsOrderRequest
	qs := mux.Vars(r)
	fmt.Println(qs["id"])
	req := payload.GetItemsOrderRequest{
		ID: qs["id"],
	}
	return req, nil
}

func DecodeCreateOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func DecodeChangeOrderStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ChangeOrderStatusRequest{
		Data: payload.ChangeOrderStatusAtributes{
			ID: qs["id"],
		},
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
