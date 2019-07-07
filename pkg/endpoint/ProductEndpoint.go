package endpoint

import (
	"context"
	"errors"
	"fmt"

	models "qasir-supplier/inventory/models"
	payload "qasir-supplier/inventory/pkg/request/payload" // part of transport
	"qasir-supplier/inventory/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetProductsEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, err := srv.GetProducts()
		if err != nil {
			return payload.GetProductsResponse{nil, err.Error()}, nil
		}
		return payload.GetProductsResponse{*d, ""}, nil
	}
}

func MakeShowProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowProductRequest)
		fmt.Println("here")
		d, err := srv.ShowProduct(req.Id)

		if err != nil {
			return payload.ShowProductResponse{d, err.Error()}, nil
		}
		return payload.ShowProductResponse{d, ""}, nil
	}
}

func MakeCreateProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateProductRequest)
		d := srv.CreateProduct(req)
		return payload.CreateProductResponse{d.Message, d.StatusCode}, nil
	}
}

func MakeUpdateProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.UpdateProductRequest)
		d := srv.UpdateProduct(req, req.Id)
		return payload.UpdateProductResponse{d.Message, d.StatusCode}, nil
	}
}

func MakeDeleteProductEndpoint(srv service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.DeleteProductRequest)
		d := srv.DeleteProduct(req.Id)
		return payload.DeleteProductResponse{d.Message, d.StatusCode}, nil
	}
}

func (e Endpoints) GetProducts(ctx context.Context) ([]models.Product, error) {
	req := payload.GetProductsRequest{}
	resp, err := e.GetProductsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(payload.GetProductsResponse)
	if getResp.Err != "" {
		return nil, errors.New(getResp.Err)
	}
	return getResp.Products, nil
}

func (e Endpoints) ShowProduct(ctx context.Context) (models.Product, error) {
	req := payload.ShowProductRequest{}
	resp, _ := e.ShowProductsEndpoint(ctx, req)
	getResp := resp.(payload.ShowProductResponse)
	return getResp.Product, nil
}

func (e Endpoints) CreateProduct(ctx context.Context, data payload.CreateProductRequest) (payload.CreateProductResponse, error) {
	req := payload.CreateProductRequest{}
	resp, err := e.CreateProductEndpoint(ctx, req)
	if err != nil {
		return payload.CreateProductResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.CreateProductResponse)
	return getResp, nil
}

func (e Endpoints) UpdateProduct(ctx context.Context, data payload.UpdateProductResponse) (payload.UpdateProductResponse, error) {
	req := payload.UpdateProductRequest{}
	resp, err := e.UpdateProductEnpoint(ctx, req)
	if err != nil {
		return payload.UpdateProductResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.UpdateProductResponse)
	return getResp, nil
}

func (e Endpoints) DeleteProduct(ctx context.Context) (payload.DeleteProductResponse, error) {
	req := payload.DeleteProductRequest{}
	resp, err := e.DeleteProductEnpoint(ctx, req)
	if err != nil {
		return payload.DeleteProductResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.DeleteProductResponse)
	return getResp, nil
}
