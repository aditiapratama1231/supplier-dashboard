package request

import (
	"context"
	"net/http"
	request "qasir-supplier/inventory/pkg/request/model"
)

func DecodeGetBrandsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.GetBrandsRequest
	return req, nil
}
