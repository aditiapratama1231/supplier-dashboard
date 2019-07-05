package request

import (
	"context"
	"net/http"
	request "qasir-supplier/inventory/pkg/request/model"
)

// In the second part we will write "decoders" for our incoming requests
func DecodeGetProductsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.GetProductsRequest
	return req, nil
}
