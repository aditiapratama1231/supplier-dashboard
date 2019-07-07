package request

type GetBrandsRequest struct{}

type GetBrandsResponse struct {
	Brands string `json:"brands"`
	Err    string `json:"err,omitempty"`
}