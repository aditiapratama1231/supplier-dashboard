package transport

import (
	"context"
	"net/http"

	"qasir-supplier/inventory/pkg/endpoint"
	request "qasir-supplier/inventory/pkg/request/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	// products handlers
	getProductsHandler := httptransport.NewServer(
		endpoints.GetProductsEndpoint,
		request.DecodeGetProductsRequest,
		request.EncodeResponse,
	)

	showProductHandler := httptransport.NewServer(
		endpoints.ShowProductsEndpoint,
		request.DecodeShowProductRequest,
		request.EncodeResponse,
	)

	createProductHandler := httptransport.NewServer(
		endpoints.CreateProductEndpoint,
		request.DecodeCreateProductRequest,
		request.EncodeResponse,
	)

	updateProductHandler := httptransport.NewServer(
		endpoints.UpdateProductEnpoint,
		request.DecodeUpdateProductRequest,
		request.EncodeResponse,
	)

	deleteProductHandler := httptransport.NewServer(
		endpoints.DeleteProductEnpoint,
		request.DecodeDeleteProductRequest,
		request.EncodeResponse,
	)

	//brands handlers
	getBrandsHandler := httptransport.NewServer(
		endpoints.GetBrandsEndpoint,
		request.DecodeGetBrandsRequest,
		request.EncodeResponse,
	)

	// products client endpoints
	r.Handle("/products", getProductsHandler).Methods("GET")
	r.Handle("/products/create", createProductHandler).Methods("POST")
	r.Handle("/products/{id}", showProductHandler).Methods("GET")
	r.Handle("/products/{id}/update", updateProductHandler).Methods("PATCH")
	r.Handle("/products/{id}/delete", deleteProductHandler).Methods("DELETE")

	// brands client endpoints
	r.Handle("/brands", getBrandsHandler).Methods("GET")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
