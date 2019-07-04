package server

import (
	"context"
	"net/http"

	"qasir-supplier/inventory/pkg/endpoint"

	transport "qasir-supplier/inventory/pkg/transport/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	// products api
	r.Methods("GET").Path("/products").Handler(httptransport.NewServer(
		endpoints.GetProductsEndpoint,
		transport.DecodeGetProductsRequest,
		transport.EncodeResponse,
	))

	//brands api
	r.Methods("GET").Path("/brands").Handler(httptransport.NewServer(
		endpoints.GetBrandsEndpoint,
		transport.DecodeGetBrandsRequest,
		transport.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
