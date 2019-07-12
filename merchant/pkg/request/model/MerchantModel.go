package request

type GetMerchantsRequest struct{}

type GetMerchantsResponse struct {
	Merchants string `json:"merchants"`
	Err       string `json:"err,omitempty"`
}
