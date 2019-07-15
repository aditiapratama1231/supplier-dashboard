package request

import (
	"time"

	models "qasir-supplier/order/models"
)

type Order struct {
	VariantID      int64     `json:"variant_id"`
	SupplierID     int64     `json:"supplier_id"`
	MerchantID     int64     `json:"merchant_id"`
	OutletID       int64     `json:"outlet_id"`
	Subtotal       float64   `json:"subtotal"`
	DiscountAmount float64   `json:"discount_amount"`
	GrandTotal     float64   `json:"grand_total"`
	OrderNo        string    `json:"order_no"`
	Status         string    `json:"status"`
	HasInvoice     bool      `json:"has_invoice"`
	CanceledAt     time.Time `json:"canceled_at" db:"canceled_at"`
	ShippedAt      time.Time `json:"shipped_at" db:"shipped_at"`
	DeliveredAt    time.Time `json:"delivered_at" db:"delivered_at"`
	CompletedAt    time.Time `json:"completed_at" db:"completed_at"`
}

type GetMerchantOrdersRequest struct {
	PaginationRequest
}

type GetMerchantOrdersResponse struct {
	Message    string             `json:"message"`
	Data       []models.Order     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
	Err        bool               `json:"err"`
}

// CreateOrderRequest struct to Create New Product
type CreateOrderRequest struct {
	Data CreateOrderAtributes `json:"data,omitempty"`
}

// CreateOrderResponse struct
type CreateOrderResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
	Err        bool   `json:"error"`
}

type CreateOrderAtributes struct {
	Atributes Order `json:"atributes"`
}

// ChangeOrderStatusRequest struct to Create New Product
type ChangeOrderStatusRequest struct {
	Data ChangeOrderStatusAtributes `json:"data,omitempty"`
}

// ChangeOrderStatusResponse struct
type ChangeOrderStatusResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
	Err        bool   `json:"error"`
}

type ChangeOrderStatusAtributes struct {
	ID        string            `json:"id"`
	Atributes ChangeOrderStatus `json:"atributes"`
}

type ChangeOrderStatus struct {
	Status      string    `json:"status"`
	HasInvoice  bool      `json:"has_invoice"`
	CanceledAt  time.Time `json:"canceled_at" db:"canceled_at"`
	ShippedAt   time.Time `json:"shipped_at" db:"shipped_at"`
	DeliveredAt time.Time `json:"delivered_at" db:"delivered_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
}
