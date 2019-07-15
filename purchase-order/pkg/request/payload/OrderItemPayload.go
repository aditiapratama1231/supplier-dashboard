package request

import (
	models "qasir-supplier/order/models"
)

type OrderItem struct {
	VariantID     int64   `json:"variant_id"`
	OrderID       int64   `json:"order_id"`
	Qty           float64 `json:"qty"`
	QtyInvoice    float64 `json:"qty_invoice"`
	QtyShipped    float64 `json:"qty_shipped"`
	QtyCanceled   float64 `json:"qty_canceled"`
	QtyRefunded   float64 `json:"qty_refunded"`
	PriceSell     float64 `json:"price_sell"`
	PriceSellUnit float64 `json:"price_sellUnit"`
	PriceBase     float64 `json:"price_base"`
	PriceBaseUnit float64 `json:"price_base_unit"`
	Sku           string  `json:"sku"`
	Name          string  `json:"name"`
}

type GetItemsOrderRequest struct {
	ID string `json:"id"`
}

type GetItemsOrderResponse struct {
	Message string       `json:"message"`
	Data    models.Order `json:"data,omitempty"`
	Err     bool         `json:"err"`
}
