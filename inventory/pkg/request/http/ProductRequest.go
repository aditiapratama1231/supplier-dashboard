package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	payload "qasir-supplier/inventory/pkg/request/payload"
)

// In the second part we will write "decoders" for our incoming requests
func DecodeGetProductsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetProductsRequest
	return req, nil
}

func DecodeShowProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ShowProductRequest{
		Id: qs["id"],
	}
	return req, nil
}

func DecodeCreateProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeUpdateProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.UpdateProductRequest{
		Id: qs["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeDeleteProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.DeleteProductRequest{
		Id: qs["id"],
	}
	return req, nil
}
