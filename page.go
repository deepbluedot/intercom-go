package intercom

// PageParams determine paging information to and from the API
type PageParams struct {
	Page          int64       `json:"page" url:"page,omitempty"`
	StartingAfter string      `json:"starting_after,omitempty"`
	PerPage       int64       `json:"per_page" url:"per_page,omitempty"`
	TotalPages    int64       `json:"total_pages" url:"-"`
	Next          *PageParams `json:"next,omitempty"`
}
