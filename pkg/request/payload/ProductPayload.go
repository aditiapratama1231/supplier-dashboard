package request

import (
	models "qasir-supplier/inventory/models"
)

type Product struct {
	Name        string `json:"product_name"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
	HasText     int8   `json:"has_text"`
	Image       string `json:"image"`
	CategoryId  int64  `json:"category_id"`
	MerchantId  int64  `json:"merchant_id"`
	BrandId     int64  `json:"brand_id"`
	IsPublished int8   `json:"is_published"`
}

type GetProductsRequest struct{}

type GetProductsResponse struct {
	Products []models.Product `json:"products"`
	Err      string           `json:"err,omitempty"`
}

// Show Product
type ShowProductRequest struct {
	Id string `"json:id"`
}

type ShowProductResponse struct {
	Product models.Product `json:"product,omitempty"`
	Err     string         `json:"err,omitempty"`
}

// Create New Product
type CreateProductRequest struct {
	Product
}

type CreateProductResponse struct {
	Message    string `"json:message"`
	StatusCode int32  `"json:status_code"`
}

// Update Product
type UpdateProductRequest struct {
	Id string `"json:id"`
	Product
}

type UpdateProductResponse struct {
	Message    string `"json:message"`
	StatusCode int32  `"json:status_code"`
}

type DeleteProductRequest struct {
	Id string `"json:id"`
}

type DeleteProductResponse struct {
	Message    string `"json:message"`
	StatusCode int32  `"json:status_code"`
}
