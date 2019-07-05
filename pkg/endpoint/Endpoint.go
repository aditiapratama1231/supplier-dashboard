package endpoint

import "github.com/go-kit/kit/endpoint"

// Register our endpoints here
type Endpoints struct {
	GetProductsEndpoint endpoint.Endpoint
	GetBrandsEndpoint   endpoint.Endpoint
}
