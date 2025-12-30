package dto

type Sort struct {
	ColId string `json:"col_id"`
	Sort  string `json:"sort"`
}

type Filter struct {
	Type       string `json:"type"`
	From       string `json:"from"`
	To         string `json:"to"`
	FilterType string `json:"filter_type"`
}

type DynamicFilter struct {
	Sort   []Sort            `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

type Pagination[T any] struct {
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	PageSize   int   `json:"page_size"`
	HasPrev    bool  `json:"has_prev"`
	HasNext    bool  `json:"has_next"`
	Data       *[]T  `json:"data"`
}

type QueryParams struct {
	Search         string            `json:"search"`
	Sorts          []Sort            `json:"sorts"`
	Filters        map[string]Filter `json:"filters"`
	DynamicFilters []DynamicFilter   `json:"dynamic_filters"`
	Pagination     Pagination[any]   `json:"pagination"`
}

type PaginationInput struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type PaginationResultWithFilter struct {
	PaginationInput
	DynamicFilter
}

func (p *PaginationResultWithFilter) GetOffsetLimit() int {
	return (p.GetPageNumber() - 1) * p.GetPageSize()
}

func (p *PaginationResultWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		return 10
	} else {
		return p.PageSize
	}
}

func (p *PaginationResultWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		return 1
	} else {
		return p.PageNumber
	}
}

func (p *PaginationResultWithFilter) IsPaginationProvided() bool {
	return p.PageNumber > 0 && p.PageSize > 0
}
