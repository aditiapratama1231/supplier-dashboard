package endpoint

import "github.com/go-kit/kit/endpoint"

// Register our endpoints here
type Endpoints struct {
	GetProductsEndpoint   endpoint.Endpoint
	ShowProductsEndpoint  endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
	UpdateProductEnpoint  endpoint.Endpoint
	DeleteProductEnpoint  endpoint.Endpoint

	GetBrandsEndpoint endpoint.Endpoint
}
