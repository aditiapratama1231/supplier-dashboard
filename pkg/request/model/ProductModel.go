package request

type GetProductsRequest struct{}

type GetProductsResponse struct {
	Products string `json:"products"`
	Err      string `json:"err,omitempty"`
}
