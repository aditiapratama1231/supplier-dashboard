package request

import (
	"context"
	"net/http"
	payload "qasir-supplier/inventory/pkg/request/payload"
)

func DecodeGetBrandsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetBrandsRequest
	return req, nil
}
