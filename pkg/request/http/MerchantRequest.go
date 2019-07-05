package request

import (
	"context"
	"net/http"
	request "qasir-supplier/merchant/pkg/request/model"
)

// In the second part we will write "decoders" for our incoming requests
func DecodeGetMerchantsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.GetMerchantsRequest
	return req, nil
}
