package transport

import (
	"context"
	"net/http"

	"qasir-supplier/merchant/pkg/endpoint"
	request "qasir-supplier/merchant/pkg/request/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	//merchants handlers
	getMerchantsHandler := httptransport.NewServer(
		endpoints.GetMerchantsEndpoint,
		request.DecodeGetMerchantsRequest,
		request.EncodeResponse,
	)

	// brands client endpoints
	r.Handle("/merchants", getMerchantsHandler).Methods("GET")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
