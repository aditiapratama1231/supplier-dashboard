package transport

import (
	"context"
	"net/http"

	"qasir-supplier/order/pkg/endpoint"
	request "qasir-supplier/order/pkg/request/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	// products handlers
	getMerchantOrders := httptransport.NewServer(
		endpoints.GetMerchantOrdersEndpoint,
		request.DecodeGetMerchantOrdersRequest,
		request.EncodeResponse,
	)

	showOrderDetail := httptransport.NewServer(
		endpoints.ShowOrderDetailEndpoint,
		request.DecodeShowOrderDetailRequest,
		request.EncodeResponse,
	)

	createOrderHandler := httptransport.NewServer(
		endpoints.CreateOrderEndpoint,
		request.DecodeCreateOrderRequest,
		request.EncodeResponse,
	)

	changeOrderStatusHandler := httptransport.NewServer(
		endpoints.ChangeOrderStatusEndpoint,
		request.DecodeChangeOrderStatusRequest,
		request.EncodeResponse,
	)

	// orders client endpoints
	orders := r.PathPrefix("/orders").Subrouter().StrictSlash(true)
	orders.Handle("/", getMerchantOrders).Methods("GET")
	orders.Handle("/{id}", showOrderDetail).Methods("GET")
	orders.Handle("/create", createOrderHandler).Methods("POST")
	orders.Handle("/{id}/change", changeOrderStatusHandler).Methods("PATCH")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
