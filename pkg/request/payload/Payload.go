package request

type PaginationRequest struct {
	Include []string `json:"include,omitempty"`
	Page    uint32   `json:"offset,omitempty"`
	Limit   uint32   `json:"limit,omitempty"`
	// Filter  FilterProductRequest `json:"filter,omitempty"`
	// Search  Search               `json:"search,omitempty"`
}

// PaginationResponse struct contain variable to response query
type PaginationResponse struct {
	CurrentPage uint32 `json:"current_page"`
	From        uint32 `json:"from"`
	PerPage     uint32 `json:"per_page"`
	Total       uint32 `json:"total"`
	To          uint32 `json:"to"`
}
