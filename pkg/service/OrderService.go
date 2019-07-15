package service

import (
	models "qasir-supplier/order/models"
	payload "qasir-supplier/order/pkg/request/payload"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderService interface {
	GetMerchantOrders(payload.GetMerchantOrdersRequest) payload.GetMerchantOrdersResponse
	ShowOrderDetail(payload.GetItemsOrderRequest) payload.GetItemsOrderResponse
	CreateOrder(payload.CreateOrderRequest) payload.CreateOrderResponse
	ChangeOrderStatus(data payload.ChangeOrderStatusRequest) payload.ChangeOrderStatusResponse
}

type orderService struct{}

var query *gorm.DB

func NewOrderService(db *gorm.DB) OrderService {
	query = db
	return orderService{}
}

func (orderService) GetMerchantOrders(data payload.GetMerchantOrdersRequest) payload.GetMerchantOrdersResponse {
	var (
		orders []models.Order
		count  uint32
	)

	offset := (data.PaginationRequest.Page - 1) * data.PaginationRequest.Limit

	query.Find(&orders).Count(&count)

	return payload.GetMerchantOrdersResponse{
		Message: "Orders Successfully Loaded",
		Data:    orders,
		Pagination: payload.PaginationResponse{
			CurrentPage: data.PaginationRequest.Page,
			From:        offset + 1,
			To:          data.PaginationRequest.Page * data.PaginationRequest.Limit,
			PerPage:     data.PaginationRequest.Limit,
			Total:       count,
		},
		Err: false,
	}
}

func (orderService) ShowOrderDetail(data payload.GetItemsOrderRequest) payload.GetItemsOrderResponse {
	var order models.Order
	if query.Find(&order, data.ID).RecordNotFound() {
		return payload.GetItemsOrderResponse{
			Message: "Product Not Found",
			Err:     true,
		}
	}

	query.Model(order).Related(&order.OrderItems)

	return payload.GetItemsOrderResponse{
		Message: "Order Detail Successfully Loaded",
		Data:    order,
		Err:     false,
	}
}

func (orderService) CreateOrder(data payload.CreateOrderRequest) payload.CreateOrderResponse {
	order := models.Order{
		VariantID:      data.Data.Atributes.VariantID,
		SupplierID:     data.Data.Atributes.SupplierID,
		MerchantID:     data.Data.Atributes.MerchantID,
		OutletID:       data.Data.Atributes.OutletID,
		Subtotal:       data.Data.Atributes.Subtotal,
		DiscountAmount: data.Data.Atributes.DiscountAmount,
		GrandTotal:     data.Data.Atributes.GrandTotal,
		OrderNo:        data.Data.Atributes.OrderNo,
		Status:         data.Data.Atributes.Status,
		HasInvoice:     data.Data.Atributes.HasInvoice,
	}

	err := query.Create(&order).Error
	if err != nil {
		return payload.CreateOrderResponse{
			Message:    err.Error(),
			Err:        false,
			StatusCode: 500,
		}
	}
	return payload.CreateOrderResponse{
		Message:    "Success created order",
		Err:        false,
		StatusCode: 201,
	}
}

func (orderService) ChangeOrderStatus(data payload.ChangeOrderStatusRequest) payload.ChangeOrderStatusResponse {
	var order models.Order
	if query.Find(&order, data.Data.ID).RecordNotFound() {
		return payload.ChangeOrderStatusResponse{
			Message:    "Order Not Found",
			StatusCode: 404,
			Err:        true,
		}
	}
	canceledAt, _ := time.Parse(time.RFC3339, "0001-01-01 00:00:00 +0000 UTC")
	shippedAt, _ := time.Parse(time.RFC3339, "0001-01-01 00:00:00 +0000 UTC")
	deliveredAt, _ := time.Parse(time.RFC3339, "0001-01-01 00:00:00 +0000 UTC")
	completedAt, _ := time.Parse(time.RFC3339, "0001-01-01 00:00:00 +0000 UTC")
	if data.Data.Atributes.CanceledAt != canceledAt {
		order.CanceledAt = data.Data.Atributes.CanceledAt
	}
	if data.Data.Atributes.DeliveredAt != deliveredAt {
		order.DeliveredAt = data.Data.Atributes.DeliveredAt
	}
	if data.Data.Atributes.ShippedAt != shippedAt {
		order.ShippedAt = data.Data.Atributes.ShippedAt
	}
	if data.Data.Atributes.CompletedAt != completedAt {
		order.CompletedAt = data.Data.Atributes.CompletedAt
	}
	if data.Data.Atributes.Status != "" {
		order.Status = data.Data.Atributes.Status
	}

	order.HasInvoice = data.Data.Atributes.HasInvoice

	err := query.Save(&order).Error
	if err != nil {
		return payload.ChangeOrderStatusResponse{
			Message:    err.Error(),
			StatusCode: 500,
			Err:        true,
		}
	}
	return payload.ChangeOrderStatusResponse{
		Message:    "Success change order status to be " + data.Data.Atributes.Status,
		Err:        false,
		StatusCode: 202,
	}
}
