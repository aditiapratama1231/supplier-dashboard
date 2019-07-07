package service

import (
	db "qasir-supplier/inventory/database"
	models "qasir-supplier/inventory/models"
	payload "qasir-supplier/inventory/pkg/request/payload" // part of transport
)

type ProductService interface {
	GetProducts() (*[]models.Product, error)
	ShowProduct(string) (models.Product, error)
	CreateProduct(payload.CreateProductRequest) payload.CreateProductResponse
	UpdateProduct(payload.UpdateProductRequest, string) payload.UpdateProductResponse
	DeleteProduct(string) payload.DeleteProductResponse
}

type productService struct{}

func NewProdutService() ProductService {
	return productService{}
}

// Get will return today's date
func (productService) GetProducts() (*[]models.Product, error) {
	var products []models.Product

	query := db.DBInit()
	query.Find(&products)

	defer db.DBInit().Close()

	return &products, nil
}

func (productService) ShowProduct(id string) (models.Product, error) {
	var product models.Product

	query := db.DBInit()
	query.Find(&product, id)

	defer db.DBInit().Close()

	return product, nil
}

func (productService) CreateProduct(data payload.CreateProductRequest) payload.CreateProductResponse {
	product := models.Product{
		Name:        data.Name,
		ProductCode: data.ProductCode,
		Description: data.Description,
		CategoryID:  data.CategoryId,
		MerchantID:  data.MerchantId,
		BrandID:     data.MerchantId,
		Image:       data.Image,
		HasText:     data.HasText,
		IsPublished: data.IsPublished,
	}
	err := models.InsertProduct(db.DBInit(), &product)
	if err != nil {
		return payload.CreateProductResponse{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	defer db.DBInit().Close()
	return payload.CreateProductResponse{
		Message:    "Product Successfully Created",
		StatusCode: 201,
	}
}

func (productService) UpdateProduct(data payload.UpdateProductRequest, id string) payload.UpdateProductResponse {
	var product models.Product
	query := db.DBInit()
	if query.Find(&product, id).RecordNotFound() {
		return payload.UpdateProductResponse{
			Message:    "Product Not Found",
			StatusCode: 404,
		}
	}
	product.Name = data.Product.Name
	product.ProductCode = data.Product.ProductCode
	product.Description = data.Product.Description
	product.CategoryID = data.Product.CategoryId
	product.MerchantID = data.Product.MerchantId
	product.BrandID = data.Product.BrandId
	product.Image = data.Product.Image
	product.HasText = data.Product.HasText
	product.IsPublished = data.Product.IsPublished

	query.Save(&product)

	return payload.UpdateProductResponse{
		Message:    "Product Successfully updated",
		StatusCode: 200,
	}
}

func (productService) DeleteProduct(id string) payload.DeleteProductResponse {
	var product models.Product
	query := db.DBInit()
	if query.Find(&product, id).RecordNotFound() {
		return payload.DeleteProductResponse{
			Message:    "Product Not Found",
			StatusCode: 404,
		}
	}
	query.Delete(&product)

	return payload.DeleteProductResponse{
		Message:    "Product Successfully Deleted",
		StatusCode: 200,
	}
}
