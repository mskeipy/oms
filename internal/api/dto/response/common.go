package response

type CommonResponse struct {
	Data  interface{}    `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int          `json:"code"`
	Errors  []FieldError `json:"errors,omitempty"`
	Message string       `json:"message,omitempty"`
}

type Pagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
	TotalPage int   `json:"total_page"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func NewPagination(page, pageSize int, total int64) Pagination {
	totalPage := int((total + int64(pageSize) - 1) / int64(pageSize))
	return Pagination{
		Page:      page,
		PageSize:  pageSize,
		TotalRows: total,
		TotalPage: totalPage,
	}
}
