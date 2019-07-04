package transport

import (
	"context"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type GetProductsRequest struct{}

type GetProductsResponse struct {
	Products string `json:"products"`
	Err      string `json:"err,omitempty"`
}

// In the second part we will write "decoders" for our incoming requests
func DecodeGetProductsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetProductsRequest
	return req, nil
}
