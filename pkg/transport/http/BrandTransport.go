package transport

import (
	"context"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type GetBrandsRequest struct{}

type GetBrandsResponse struct {
	Brands string `json:"brands"`
	Err    string `json:"err,omitempty"`
}

// In the second part we will write "decoders" for our incoming requests
func DecodeGetBrandsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetBrandsRequest
	return req, nil
}
