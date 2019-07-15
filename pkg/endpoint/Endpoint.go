package endpoint

import "github.com/go-kit/kit/endpoint"

// Register our endpoints here
type Endpoints struct {
	GetMerchantOrdersEndpoint endpoint.Endpoint
	ShowOrderDetailEndpoint   endpoint.Endpoint
	CreateOrderEndpoint       endpoint.Endpoint

	ChangeOrderStatusEndpoint endpoint.Endpoint
}
